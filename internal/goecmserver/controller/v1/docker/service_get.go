package docker

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (d *DockerController) ServiceGet(c *gin.Context) {
	log.Debug("Service get function called.")

	data, err := d.srv.Docker().ServiceGet(c, metav1.ListOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, data)
}
