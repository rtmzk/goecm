package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/pkg/code"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/core"
)

func Resize(c *gin.Context) {
	var width = c.Query("w")
	var height = c.Query("h")
	var containerId = c.Query("cid")

	cli, _ := docker.NewDockerClient()
	if err := cli.Resize(c, containerId, width, height); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "无法调整窗口大小"), nil)
		return
	}
}
