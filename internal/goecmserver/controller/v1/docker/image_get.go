package docker

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
	"strings"
)

func (d *DockerController) ImageGet(c *gin.Context) {
	log.L(c).Debug("get image function called.")
	var hosts []string
	//port , _ := c.Params.Get("port")
	host := c.Query("hosts")
	hosts = strings.Split(host, ",")
	data, err := d.srv.Docker().ImageGet(c, hosts, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, data)
}
