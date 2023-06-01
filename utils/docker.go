package utils

import (
	"context"
	"github.com/docker/docker/client"
	"go-ecm/pkg/log"
)

func NewDockerClient(ctx context.Context, opts ...client.Opt) *client.Client {
	log.L(ctx).Debug("try get docker client")
	cli, err := client.NewClientWithOpts(opts...)
	if err != nil {
		log.L(ctx).Error("获取docker客户端失败")
		return nil
	}

	return cli
}
