package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/core"
	"io"
)

func ExportImage(c *gin.Context) {
	var data []string
	if err := c.ShouldBindJSON(&data); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	var cli, _ = docker.NewDockerClient()

	resp, err := cli.ExportImage(context.Background(), data)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	io.Copy(c.Writer, resp)
}
