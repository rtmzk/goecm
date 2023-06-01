package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"strconv"
)

func (cli *dockerClient) Resize(ctx context.Context, containerid, w, h string) error {
	c := cli.cli
	width, _ := strconv.ParseUint(w, 10, 0)
	height, _ := strconv.ParseUint(h, 10, 0)
	return c.ContainerResize(ctx, containerid, types.ResizeOptions{
		Width:  uint(width),
		Height: uint(height),
	})
}
