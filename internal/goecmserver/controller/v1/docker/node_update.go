package docker

import (
	"github.com/docker/docker/api/types/swarm"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (d *DockerController) NodeUpdate(c *gin.Context) {
	log.L(c).Debug("Node update function called.")

	var node swarm.Node
	err := c.ShouldBindJSON(&node)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "客户端传参错误"), nil)
		return
	}

	err = d.srv.Docker().NodeUpdate(c, &node, metav1.UpdateOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
