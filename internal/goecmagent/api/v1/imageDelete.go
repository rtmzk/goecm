package v1

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"go-ecm/internal/goecmagent/utils"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/core"
)

func DeleteImage(c *gin.Context) {
	var data []string
	if err := c.ShouldBindJSON(&data); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	var cli, _ = docker.NewDockerClient()

	for _, id := range data {
		cli.DeleteImage(context.Background(), id, types.ImageRemoveOptions{})
	}

	utils.ManualReport(c, cli)
	core.WriteResponse(c, nil, map[string]bool{"ack": true})
}
