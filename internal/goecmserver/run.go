package goecmserver

import "go-ecm/internal/goecmserver/config"

func Run(cfg *config.Config) error {
	server, err := createGOECMServer(cfg)
	if err != nil {
		return err
	}

	return server.PrapareRun().Run()
}
