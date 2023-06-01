package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

func (cli *dockerClient) Exec(ctx context.Context, containerId, command string) (types.HijackedResponse, error) {
	c := cli.cli
	if command == "" {
		command = "/bin/bash"
	}
	_ = cli.Resize(ctx, containerId, "200", "30")
	stream, err := c.ContainerExecCreate(ctx, containerId, types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{command},
		Tty:          true,
		Env:          []string{"COLUMNS=200", "ROWS=30"},
	})
	if err != nil {
		return types.HijackedResponse{}, err
	}
	return c.ContainerExecAttach(ctx, stream.ID, types.ExecStartCheck{Detach: false, Tty: true})
}
