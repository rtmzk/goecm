package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"go-ecm/pkg/log"
)

func (cli *dockerClient) GetContainers(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	var c = cli.cli
	data, err := c.ContainerList(ctx, options)
	if err != nil {
		return nil, err
	}
	log.Debugf("docker/GetContainers: %#v", &data)

	return data, nil
}

func (cli *dockerClient) GetContainerInspect(id string) (*types.ContainerJSON, error) {
	c := cli.cli
	data, err := c.ContainerInspect(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (cli *dockerClient) DeleteContainer(ctx context.Context, containerId string, option types.ContainerRemoveOptions) error {
	c := cli.cli
	return c.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{Force: true})
}
