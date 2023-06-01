package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

func (c *dockerClient) NetworkList(ctx context.Context, options types.NetworkListOptions) ([]types.NetworkResource, error) {
	var cli = c.cli
	return cli.NetworkList(ctx, options)
}

func (c *dockerClient) NetworkCreate(ctx context.Context, name string, create types.NetworkCreate) error {
	var cli = c.cli

	_, err := cli.NetworkCreate(ctx, name, create)
	if err != nil {
		return err
	}
	return nil
}
