package docker

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (d *DockerController) NodeInspect(c *gin.Context) {
	log.L(c).Debug("Node inspect function called.")

	data, err := d.srv.Docker().NodeInspect(c, c.Param("id"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, data)
}
