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

func (d *DockerController) ServiceUpdate(c *gin.Context) {
	log.L(c).Debug("Service update function called.")

	var service swarm.Service
	err := c.ShouldBindJSON(&service)
	if err != nil {
		log.L(c).Errorf("客户端body传参错误. Error: %s", err.Error())
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "客户端传参错误"), nil)
	}
	err = d.srv.Docker().ServiceUpdate(c, &service, metav1.UpdateOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
