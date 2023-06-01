package docker

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (d *DockerController) NodeDelete(c *gin.Context) {
	log.L(c).Debug("Node Delete function called.")

	nodeId := c.Param("id")
	err := d.srv.Docker().NodeDelete(c, nodeId, metav1.DeleteOptions{})
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrNodeDelete, "删除节点失败. Error: %s", err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
