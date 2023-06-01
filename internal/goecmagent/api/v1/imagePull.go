package v1

import (
	"github.com/docker/docker/api/types"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmagent/utils"
	"go-ecm/internal/pkg/code"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/core"
	"io"
)

func ImagePull(c *gin.Context) {
	image := c.Query("name")
	if image == "" {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "镜像名称为空"), nil)
		return
	}
	cli, _ := docker.NewDockerClient()
	resp, err := cli.PullImage(c, image, types.ImagePullOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrImagePullWithConnectionErr, "无法拉取镜像"), nil)
		return
	}
	defer resp.Close()

	io.Copy(c.Writer, resp)
	utils.ManualReport(c, cli)
}
