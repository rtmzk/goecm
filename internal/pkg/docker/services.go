package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"go-ecm/pkg/log"
)

func (cli *dockerClient) GetServices(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error) {
	c := cli.cli

	data, err := c.ServiceList(ctx, options)
	if err != nil {
		log.Warnf("Failed get service list. error: %s", err)
		return nil, err
	}
	log.Debugf("docker/GetServices: %#v", &data)

	return data, nil
}
