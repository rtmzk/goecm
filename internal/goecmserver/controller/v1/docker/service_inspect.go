package docker

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (d *DockerController) ServiceInspect(c *gin.Context) {
	log.Debug("Service inspect function called.")
	data, err := d.srv.Docker().ServiceInspect(c, c.Param("id"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, data)
}
