package goecmagent

import "go-ecm/internal/goecmagent/config"

func Run(cfg *config.Config) error {
	server, err := createGOECMAgent(cfg)
	if err != nil {
		return err
	}

	return server.PrapareRun().Run()
}
