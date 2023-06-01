package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmserver/constand"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/resource"
	"go-ecm/internal/goecmserver/util"
	"go-ecm/internal/pkg/code"
	"go-ecm/internal/pkg/docker"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/utils"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

func (a *AgentController) DeployTask(c *gin.Context) {
	var deploySpec v1.DeploySpec

	err := c.ShouldBindJSON(&deploySpec)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "Error bind request body"), nil)
		return
	}

	core.WriteResponse(c, nil, map[string]bool{"task": true})
	err = a.deploy(&deploySpec)
	if err != nil {
		updateDeployProcess("", "部署错误: "+err.Error())
		_ = a.srv.Agent().UpdateInitStatus(0)
	}
}

func (a *AgentController) GetTaskProcess(c *gin.Context) {
	core.WriteResponse(c, nil, constand.DeployProcess)
}

func (a *AgentController) deploy(deploySpec *v1.DeploySpec) error {
	var agents []v1.Agent
	var uploadWg sync.WaitGroup
	var deployWg sync.WaitGroup

	errs := flushTemplate(deploySpec)
	if len(errs) != 0 {
		return errors.New("刷新模板时发生错误")
	}
	agents, err := a.srv.Agent().GetAgents(context.Background(), metav1.GetOptions{})
	if err != nil {
		return err
	}
	uploadWg.Add(len(agents))
	deployWg.Add(len(agents))

	//并行上传文件,后续的操作要等文件上传完之后再进行
	updateDeployProcess("5%", "正在传输镜像包")
	for _, v := range agents {
		if v.Deployed {
			uploadWg.Done()
			continue
		}
		go func(a v1.Agent) {
			defer uploadWg.Done()
			var addr = fmt.Sprintf("%s:%s", a.AgentAddr, strconv.Itoa(a.AgentPort))
			_ = util.MultiPartUploadHelper("/home/ecm/Install_Packages/Image_ecm_mainregion_"+constand.MainPackageVersion+".tar", addr, "/opt/", 100)
			_ = util.MultiPartUploadHelper("/home/ecm/Install_Packages/portainer_v2.0.tar", addr, "/opt/", 20)
			_ = util.MultiPartUploadHelper("/home/ecm/Install_Packages/apm_mainregion_"+constand.MainPackageVersion+".tar", addr, "/opt/", 20)
			return
		}(v)
	}
	uploadWg.Wait()
	updateDeployProcess("40%", "正在导入镜像包")

	//部署步骤在agent中实现
	for _, v := range agents {
		if v.Deployed {
			deployWg.Done()
			continue
		}
		go a.deployStepByStep(&deployWg, v, deploySpec)
	}
	deployWg.Wait()

	updateDeployProcess("80%", "正在更新swarm节点")
	_ = updateNodeLabel(deploySpec)
	createOverlayNetwork(deploySpec)
	// 启动服务
	updateDeployProcess("97%", "正在启动服务")
	initStack(deploySpec)

	time.Sleep(time.Second * 30)
	updateDeployProcess("100%", "部署完成")
	return nil
}

func flushTemplate(spec *v1.DeploySpec) []error {
	var tmpls []*template.Template
	var errs []error
	var funcMap = util.NewFuncMap()

	tmpls = append(tmpls,
		template.Must(template.New("middleware-master-cluster.yml").Parse(string(resource.MiddlewareClusterYML))),
		template.Must(template.New("middleware-master-standalone.yml").Parse(string(resource.MiddlewareStandaloneYML))),
		template.Must(template.New("envfile.env").Parse(string(resource.EnvFile))),
		template.Must(template.New("docker-compose-mainregion.yml").Funcs(funcMap).Parse(string(resource.MainRegionYML))),
		template.Must(template.New("docker-compose-inbiz.yml").Parse(string(resource.InbizYML))),
		template.Must(template.New("docker-compose-apm.yml").Parse(string(resource.APMYML))),
	)

	for _, t := range tmpls {
		var buf bytes.Buffer
		if err := t.Execute(&buf, spec); err != nil {
			errs = append(errs, err)
			continue
		}
		content, _ := ioutil.ReadAll(&buf)
		if utils.IsExist("/opt/" + t.Name()) {
			os.RemoveAll("/opt" + t.Name())
		}
		utils.SaveFile("/opt/"+t.Name(), string(content))
	}
	return errs
}

func updateNodeLabel(spec *v1.DeploySpec) error {
	var appNode []swarm.Node
	var middlewareAddr = make(map[string][]string)
	var mode = spec.MiddlewareDeploySpec.Mode
	middlewareAddr["es"] = utils.StrSplit(spec.MiddlewareDeploySpec.Elasticsearch.ElasticsearchHosts, ",")
	middlewareAddr["redis"] = utils.StrSplit(spec.MiddlewareDeploySpec.Redis.RedisHosts, ",")
	middlewareAddr["mysql"] = utils.StrSplit(spec.MiddlewareDeploySpec.Database.DBHosts, ",")
	middlewareAddr["rabbitmq"] = utils.StrSplit(spec.MiddlewareDeploySpec.Redis.RedisHosts, ",")
	cli, err := docker.NewDockerClient()
	if err != nil {
		return err
	}

	nodes, err := cli.GetNodes(types.NodeListOptions{})
	if err != nil {
		return err
	}

	for _, n := range nodes {
		if strings.Contains(spec.Hosts, n.Status.Addr) {
			appNode = append(appNode, n)
		}
	}

	for _, n := range appNode {
		n.Spec.Labels = map[string]string{
			"nodetype": "InDrive",
		}
		n.Spec.Role = "manager"
		_ = cli.UpdateNode(&n)
	}

	var middlewareNodes = findMiddlewareNode(nodes, middlewareAddr)
	for k, mn := range middlewareNodes {
		var startIndex = 1
		var mysqlIndex = 0
		for _, n := range mn {
			node, _ := cli.NodeInspect(context.Background(), n.ID)
			n = node
			if mode == "standalone" {
				n.Spec.Labels["nodelabels"] = "Middleware"
				err = cli.UpdateNode(&n)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			if k == "mysql" {
				n.Spec.Labels[k+"labels"] = k + strconv.Itoa(mysqlIndex)
				_ = cli.UpdateNode(&n)
			} else {
				n.Spec.Labels[k+"labels"] = k + strconv.Itoa(startIndex)
				_ = cli.UpdateNode(&n)
			}
			startIndex = startIndex + 1
			mysqlIndex = mysqlIndex + 1
		}
	}

	return nil
}

func findMiddlewareNode(nodes []swarm.Node, addrs map[string][]string) map[string][]swarm.Node {
	var middlewareNodes = make(map[string][]swarm.Node)

	// 服务对应的节点ip列表
	for key, addr := range addrs {
		for _, a := range addr {
			for _, n := range nodes {
				if n.Status.Addr == a {
					middlewareNodes[key] = append(middlewareNodes[key], n)
				}
			}
		}

	}

	return middlewareNodes
}

func (a *AgentController) deployStepByStep(wg *sync.WaitGroup, agent v1.Agent, spec *v1.DeploySpec) {
	defer wg.Done()
	var url string
	var faile []struct{}

	var mainPackageStr = fmt.Sprintf("?name=%s&loaded=%t", "Image_ecm_mainregion_"+constand.MainPackageVersion+".tar", agent.ImageLoaded)
	var apmPackageStr = fmt.Sprintf("?name=%s&loaded=%t", "apm_mainregion_"+constand.MainPackageVersion+".tar", agent.ImageLoaded)
	var portainerPackageStr = fmt.Sprintf("?name=%s&loaded=%t", "portainer_v2.0.tar", agent.ImageLoaded)

	var actions = []v1.Actions{
		{
			Action: "CheckDataPath",
			Data:   spec,
			Query:  "",
		},
		{
			Action: "LoadImage",
			Data:   nil,
			Query:  mainPackageStr,
		},
		{
			Action: "LoadImage",
			Data:   nil,
			Query:  apmPackageStr,
		},
		{
			Action: "LoadImage",
			Data:   nil,
			Query:  portainerPackageStr,
		},
	}
	var addr = fmt.Sprintf("%s:%s", agent.AgentAddr, strconv.Itoa(agent.AgentPort))
	for _, op := range actions {
		url = fmt.Sprintf("http://%s/v1/DeployCore/%s%s", addr, op.Action, op.Query)
		var cli http.Client
		var req *http.Request
		if op.Data != nil {
			body, err := json.Marshal(op.Data)
			if err != nil {
				return
			}
			req, _ = http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
		} else {
			req, _ = http.NewRequest(http.MethodPost, url, nil)
		}
		resp, _ := cli.Do(req)
		if err := ensureReadResponse(resp); err != nil {
			faile = append(faile, struct{}{})
		}
	}
	if len(faile) == 0 {
		agent.Deployed = true
		agent.ImageLoaded = true
		agent.Joined = true
		_ = a.srv.Agent().UpdateAgents(context.Background(), metav1.UpdateOptions{}, &agent)
	}
	return
}

func ensureReadResponse(response *http.Response) error {
	if response == nil {
		return errors.New("can not get response, may be step execute failed.")
	}
	defer response.Body.Close()
	var commonResp struct {
		Result string
	}
	body, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(body, &commonResp)
	if err != nil {
		return err
	}
	if commonResp.Result != "ok" {
		return errors.New("step execute failed.")
	}
	return nil
}

func createOverlayNetwork(spec *v1.DeploySpec) {
	var cli, _ = docker.NewDockerClient()
	var findFlag = false
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return
	}
	for _, n := range networks {
		if n.Name == "macrowing" {
			findFlag = true
		}
	}
	if findFlag {
		return
	}
	create := types.NetworkCreate{
		Driver:         "overlay",
		CheckDuplicate: true,
		EnableIPv6:     false,
		IPAM: &network.IPAM{
			Driver: "default",
			Config: []network.IPAMConfig{
				{
					Subnet: spec.DockerOverlayMacrowingNetwork,
				},
			},
		},
		Internal: false,
	}
	_ = cli.NetworkCreate(context.Background(), "macrowing", create)
}

func initStack(spec *v1.DeploySpec) {
	var args []string
	//binaryPath := "/hostfs/usr/bin/docker"
	binaryPath := "/usr/bin/docker"
	var middlewareCompose string
	//var mainServerCompose = []string{"/hostfs/opt/docker-compose-mainregion.yml", "/hostfs/opt/docker-compose-inbiz.yml", "/hostfs/opt/docker-compose-apm.yml", "/hostfs/opt/portainer-agent-stack.yml"}
	var mainServerCompose = make(map[string]string)
	mainServerCompose["indrive"] = "/opt/docker-compose-mainregion.yml"
	mainServerCompose["inbiz"] = "/opt/docker-compose-inbiz.yml"
	mainServerCompose["apm"] = "/opt/docker-compose-apm.yml"
	mainServerCompose["portainer"] = "/opt/portainer-agent-stack.yml"

	if spec.Mode == "standalone" {
		middlewareCompose = "middleware-master-standalone.yml"
	} else {
		middlewareCompose = "middleware-master-cluster.yml"
	}
	args = append(args, "stack", "deploy", "-c", "/opt/"+middlewareCompose, "middleware")
	middlewareDeployCmd := exec.Command(binaryPath, args...)
	_ = middlewareDeployCmd.Run()

	time.Sleep(30)
	for stackName, composePath := range mainServerCompose {
		_ = exec.Command(binaryPath, "stack", "deploy", "-c", composePath, stackName).Run()
	}
}

func updateDeployProcess(percent, message string) {
	if percent != "" {
		constand.DeployProcess.Percent = percent
	}
	constand.DeployProcess.Message = message
}
