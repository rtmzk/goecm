package goecmserver

import (
	"bytes"
	"context"
	"go-ecm/internal/goecmserver/config"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/resource"
	"go-ecm/internal/goecmserver/store"
	"go-ecm/internal/goecmserver/store/sqlite"
	"go-ecm/internal/goecmserver/util/docker"
	metav1 "go-ecm/internal/pkg/meta/v1"
	genericoptions "go-ecm/internal/pkg/options"
	genericgoecmserver "go-ecm/internal/pkg/server"
	"go-ecm/pkg/log"
	"go-ecm/pkg/shutdown"
	"go-ecm/pkg/shutdown/shutdownmanagers/posixsignal"
	"go-ecm/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	genericAPIServer *genericgoecmserver.GenericServer
}

type preparedAPIServer struct {
	*apiServer
}

type ExtraConfig struct {
	sqliteOptions      *genericoptions.SQLiteOptions
	Docker0NetPool     string
	GwbridgeNetPool    string
	MacrowingNetPool   string
	swarmAdvertiseAddr string
}

func createGOECMServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	extraConfig.complete().New()

	server := &apiServer{
		gs:               gs,
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *apiServer) PrapareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		//mysqlStore, _ := mysql.GetMySQLFactoryOr(nil)
		sqliteStore, _ := sqlite.GetSqliteFactoryOr(nil)
		if sqliteStore != nil {
			sqliteStore.Close()
		}

		s.genericAPIServer.Close()

		return nil
	}))
	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericAPIServer.Run()
}

type completedExtraConfig struct {
	*ExtraConfig
}

func (c *ExtraConfig) complete() *completedExtraConfig {
	return &completedExtraConfig{c}
}

func (c *ExtraConfig) New() error {
	storeIns, _ := sqlite.GetSqliteFactoryOr(c.sqliteOptions)
	store.SetClient(storeIns)

	//dockerIns, _ := docker.GetDockerClientOr()
	//docker.SetClient(dockerIns)

	return prepareDockerEnv(c)
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericgoecmserver.Config, lastErr error) {
	genericConfig = genericgoecmserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.Apply(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.SecureServing.Apply(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecureServing.Apply(genericConfig); lastErr != nil {
		return
	}

	return
}

func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		sqliteOptions:      cfg.SQLiteOptions,
		swarmAdvertiseAddr: strings.Split(cfg.InsecureServing.BindAddress, ":")[0],
		Docker0NetPool:     cfg.DockerNetwork.Docker0Net,
		GwbridgeNetPool:    cfg.DockerNetwork.DockerGwbridgeNet,
		MacrowingNetPool:   cfg.DockerNetwork.MacrowingOverlayNet,
	}, nil
}

func prepareDockerEnv(c *ExtraConfig) error {
	if utils.IsExist("/usr/bin/docker") {
		return nil
	}
	var err error
	var cli *client.Client
	log.Info("docker is not install. now install docker and init docker swarm ")
	if utils.IsExist("/tmp/docker.rpm") {
		_ = os.Remove("/tmp/docker.rpm")
	}
	_ = os.WriteFile("/tmp/docker.rpm", resource.DOCKER_PACKAGE, 0644)
	err = exec.Command("rpm", "-Uvh", "/tmp/docker.rpm").Run()
	if err != nil {
		log.Error("docker install failed.")
		return err
	}

	err = prepareDockerDaemonConf(c)
	if err != nil {
		return err
	}
	_ = exec.Command("systemctl", "enable", "docker", "--now").Run()
	log.Info("docker was installed successfully.")

	cli, err = docker.GetDockerClientOr()
	if cli == nil || err != nil {
		return err
	}
	docker.SetClient(cli)

	return prepareSwarmEnv(c)
}

func prepareSwarmEnv(c *ExtraConfig) error {
	swarmInitRequest := swarm.InitRequest{
		ListenAddr:       c.swarmAdvertiseAddr,
		DataPathAddr:     "",
		ForceNewCluster:  false,
		AutoLockManagers: false,
		Availability:     "active",
		AdvertiseAddr:    c.swarmAdvertiseAddr,
	}

	_, err := docker.Client().SwarmInit(context.Background(), swarmInitRequest)
	if err != nil {
		return err
	}

	var tokens v1.SwarmJoin

	spec, err := docker.Client().SwarmInspect(context.Background())
	if err != nil {
		return err
	}

	tokens.MgrJoinToken = spec.JoinTokens.Manager
	tokens.WrkJoinToken = spec.JoinTokens.Worker

	err = store.Client().Swarm().GenderToken(context.Background(), &tokens, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return prepareMacrowingNetwork(c)
}

func prepareMacrowingNetwork(c *ExtraConfig) error {
	var findFlag = false
	networks, err := docker.Client().NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return err
	}
	for _, n := range networks {
		if n.Name == "macrowing" {
			findFlag = true
		}
	}
	if findFlag {
		return nil
	}
	create := types.NetworkCreate{
		Driver:         "overlay",
		CheckDuplicate: true,
		EnableIPv6:     false,
		IPAM: &network.IPAM{
			Driver: "default",
			Config: []network.IPAMConfig{
				{
					Subnet: c.MacrowingNetPool,
				},
			},
		},
		Internal: false,
	}
	_, err = docker.Client().NetworkCreate(context.Background(), "macrowing", create)
	if err != nil {
		return err
	}

	return nil
}

func prepareDockerDaemonConf(c *ExtraConfig) error {
	var buf bytes.Buffer
	if utils.IsExist("/etc/docker/daemon.json") {
		return nil
	}
	if utils.IsNotExist("/etc/docker") {
		os.MkdirAll("/etc/docker", 0755)
	}
	var tpl = template.Must(template.New("daemon.json").Parse(string(resource.DOCKERDAMONECONF)))
	err := tpl.Execute(&buf, c)
	if err != nil {
		return err
	}

	data, _ := ioutil.ReadAll(&buf)

	return utils.SaveFile("/etc/docker/daemon.json", string(data))
}

func prepareSystemEnv() error {
	return nil
}
