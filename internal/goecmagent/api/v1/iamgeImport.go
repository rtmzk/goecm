package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/goecmagent/utils"
	"go-ecm/internal/pkg/code"
	"go-ecm/internal/pkg/docker"
	"go-ecm/pkg/core"
	"os"
)

func ImageImport(c *gin.Context) {
	path := c.Query("filePath")
	if path == "" {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "传参错误"), nil)
		return
	}
	cli, _ := docker.NewDockerClient()
	file, err := os.Open(path)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, "为找到镜像文件"), nil)
		return
	}
	defer file.Close()

	_, err = cli.LoadImage(c, file, true)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, "无法导入镜像"), nil)
		return
	}
	utils.ManualReport(c, cli)
	core.WriteResponse(c, nil, map[string]bool{"loaded": true})
}
