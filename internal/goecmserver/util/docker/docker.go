package docker

import (
	"fmt"
	"github.com/docker/docker/client"
	"sync"
)

var (
	dockerCli *client.Client
	once      sync.Once
)

// GetDockerClientOr get global docker client instance
func GetDockerClientOr() (*client.Client, error) {
	var err error
	var cli *client.Client

	once.Do(func() {
		cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

		dockerCli = cli
	})

	if dockerCli == nil || err != nil {
		return nil, fmt.Errorf("failed get docker client instance. error: %w", err)
	}

	return dockerCli, nil
}
