package docker

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (d *DockerController) ServiceDelete(c *gin.Context) {
	log.L(c).Debug("Service delete function called.")

	err := d.srv.Docker().NodeDelete(c, c.Param("id"), metav1.DeleteOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
