package swarm

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
	"go-ecm/utils"
	"os"
)

func (s SwarmController) GetScript(c *gin.Context) {
	log.L(c).Debug("get script function called.")

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=add_host_key.sh")
	//c.Header("Content-Transfer-Encoding", "application/text/plain")
	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Disposition", "attachment; filename=add_host_key.sh")

	content, _ := os.ReadFile(utils.UserHome() + "/.ssh/add_host_key.sh")

	_, err := c.Writer.Write(content)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "获取脚本文件内容失败"), nil)
	}
}
