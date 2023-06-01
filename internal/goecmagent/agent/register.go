package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/swarm"
	"go-ecm/internal/goecmagent/config"
	"go-ecm/internal/goecmagent/constand"
	"go-ecm/internal/goecmagent/static"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var (
	once          sync.Once
	retryInterval = 5 * time.Second
)

func GetRegistryOr(c *config.Config) error {
	if c == nil {
		return fmt.Errorf("failed register to goecm server.")
	}

	var err error
	if err != nil {
		return err
	}

	once.Do(func() {
		options := &registryOption{
			Token:        c.Agent.Token,
			AgentAddr:    c.InsecureServing.BindAddress,
			AgentPort:    c.InsecureServing.BindPort,
			RemoteServer: c.Agent.ServerAddr,
		}
		_ = registerRequest(options, c)
	})
	if err != nil {
		return err
	}

	return nil
}

func registerRequest(opts *registryOption, c *config.Config) error {
	var resp *http.Response

REGISTER:
	for {
		body, err := json.Marshal(opts)
		req, err := http.NewRequest(http.MethodPost, "http://"+opts.RemoteServer+"/v1/agent/register", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err = http.DefaultClient.Do(req)

		if err != nil {
			log.Warnf("failed connect to remote goecm server: %s, retry in 5 second...", opts.RemoteServer)
			time.Sleep(retryInterval)
			continue
		}
		if resp != nil && resp.StatusCode == http.StatusOK {
			break
		}
	}
	respbody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var respdata = new(registryResponse)
	err := json.Unmarshal(respbody, &respdata)
	if err != nil {
		log.Warnf("failed register to remote goecm server: %s, retry in 5 second...", opts.RemoteServer)
		time.Sleep(retryInterval)
		goto REGISTER
	}

	if respdata.Code == 110006 {
		constand.ServerAddr = opts.RemoteServer
		return nil
	}

	if respdata.Success != "ok" {
		log.Warnf("failed register to remote goecm server: %s, retry in 5 second...", opts.RemoteServer)
		time.Sleep(retryInterval)
		goto REGISTER
	}
	// 首次注册时初始化服务器环境
	envPrepare()
	setCrond()
INITSWARM:
	err = initSwarm(c, respdata.SwarmKey)
	if err != nil {
		log.Infof("failed join swarm to %s", strings.Split(c.Agent.ServerAddr, ":")[0])
		goto INITSWARM
	}

	log.Infof("success register to remote goecm server: %s", opts.RemoteServer)
	constand.ServerAddr = opts.RemoteServer
	return nil
}

func envPrepare() {
	if utils.IsExist("/tmp/env_prepare.sh") {
		_ = exec.Command("bash", "/tmp/env_prepare.sh").Run()
	}
	file, err := os.OpenFile("/tmp/env_prepare.sh", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = io.WriteString(file, string(static.ENVPREPARE))
	if err != nil {
		return
	}
	_ = exec.Command("bash", "/tmp/env_prepare.sh").Run()
	return
}

func setCrond() {
	if utils.IsExist("/tmp/crond.sh") {
		_ = exec.Command("bash", "/tmp/crond.sh").Run()
	}

	file, err := os.OpenFile("/tmp/crond.sh", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return
	}

	defer file.Close()
	_, err = io.WriteString(file, string(static.CRONDEDSCRIPT))
	if err != nil {
		return
	}
	_ = exec.Command("bash", "/tmp/env_prepare.sh").Run()
	return
}

func initSwarm(c *config.Config, joinToken string) error {
	if !utils.IsExist("/usr/bin/docker") {
		log.Info("docker is not install. now install docker and init docker swarm ")
		if utils.IsExist("/tmp/docker.rpm") {
			_ = os.Remove("/tmp/docker.rpm")
		}
		_ = os.WriteFile("/tmp/docker.rpm", static.DOCKERPKG, 0644)
		err := exec.Command("rpm", "-Uvh", "/tmp/docker.rpm").Run()
		if err != nil {
			log.Error("docker install failed.")
			return err
		}
		_ = exec.Command("systemctl", "enable", "docker", "--now").Run()
		log.Info("docker was installed successfully.")
	}
	cli, _ := docker.NewDockerClient()
	return cli.SwarmJoin(context.Background(), swarm.JoinRequest{
		AdvertiseAddr: c.InsecureServing.BindAddress,
		ListenAddr:    "0.0.0.0:2377",
		RemoteAddrs:   strings.Split(c.Agent.ServerAddr, ":"),
		JoinToken:     joinToken,
	})
}
