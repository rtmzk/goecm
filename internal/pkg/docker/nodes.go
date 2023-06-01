package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
)

func (cli *dockerClient) GetNodes(options types.NodeListOptions) ([]swarm.Node, error) {
	c := cli.cli
	data, err := c.NodeList(context.Background(), options)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cli *dockerClient) UpdateNode(node *swarm.Node) error {
	c := cli.cli
	return c.NodeUpdate(context.Background(), node.ID, node.Version, node.Spec)
}

func (cli *dockerClient) NodeInspect(ctx context.Context, nodeId string) (swarm.Node, error) {
	var c = cli.cli
	node, _, err := c.NodeInspectWithRaw(ctx, nodeId)
	return node, err
}
