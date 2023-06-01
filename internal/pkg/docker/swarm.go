package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
)

func (c *dockerClient) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error {
	var cli = c.cli
	nodes, _ := cli.NodeList(ctx, types.NodeListOptions{})
	for _, n := range nodes {
		if n.Status.Addr == req.AdvertiseAddr {
			return nil
		}
	}

	err := cli.SwarmJoin(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
