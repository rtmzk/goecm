package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"go-ecm/internal/goecmagent/config"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/log"
	"net/http"
	"time"
)

func Reporter(cfg *config.Config) {
	var c docker.DockerOperator
	for {
		cli, err := docker.NewDockerClient()
		if err != nil {
			log.Warnf("Can not get docker client instance. Did docker already installed? error: %s", err.Error())
			time.Sleep(time.Second * 30)
			continue
		}
		c = cli
		break
	}

	for {
		if err := report(c, cfg); err != nil {
			log.Errorf("Failed collect and report data to service, error: %s", err.Error())
			time.Sleep(cfg.Agent.ReportInterval * time.Second)
			continue
		}
		time.Sleep(cfg.Agent.ReportInterval * time.Second)
	}
}

func report(cli docker.DockerOperator, cfg *config.Config) error {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var reportItemsResult ReportItems
	var buf bytes.Buffer
	var err error
	var url = fmt.Sprintf("http://%s/v1/agent/report", cfg.Agent.ServerAddr)

	reportItemsResult.Host = cfg.InsecureServing.BindAddress
	//reportItemsResult.Items.ServiceItem, err = cli.GetServices(ctx, types.ServiceListOptions{})
	reportItemsResult.Items.ContainerItem, err = cli.GetContainers(ctx, types.ContainerListOptions{All: true})
	reportItemsResult.Items.ImageItem, err = cli.GetImages(ctx, types.ImageListOptions{})

	err = json.NewEncoder(&buf).Encode(reportItemsResult)
	if err != nil {
		log.Warnf("Encode collected data failed. message: %s", err.Error())
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	req.Close = true
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		resp.Body.Close()
	}

	return err
}
