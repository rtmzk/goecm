package goecmagent

import (
	"go-ecm/internal/goecmagent/agent"
	"go-ecm/internal/goecmagent/config"
	genericgoecmserver "go-ecm/internal/pkg/server"
	"go-ecm/pkg/log"
	"go-ecm/pkg/shutdown"
	"go-ecm/pkg/shutdown/shutdownmanagers/posixsignal"
)

type apiServer struct {
	gs               *shutdown.GracefulShutdown
	genericAPIServer *genericgoecmserver.GenericServer
}

type preparedAPIServer struct {
	*apiServer
}

func createGOECMAgent(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	_ = agent.GetRegistryOr(cfg)
	go agent.Heartbeat(cfg)
	go agent.Reporter(cfg)

	server := &apiServer{
		gs:               gs,
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *apiServer) PrapareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		s.genericAPIServer.Close()

		return nil
	}))
	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run() error {
	if err := s.gs.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %s", err.Error())
	}

	return s.genericAPIServer.RunAgent()
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
