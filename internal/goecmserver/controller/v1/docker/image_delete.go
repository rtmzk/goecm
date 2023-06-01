package docker

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (d *DockerController) ImageDelete(c *gin.Context) {
	log.L(c).Debug("Image delete function called.")

	var b v1.ImageDeleteRequest
	err := c.ShouldBindJSON(&b)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "Can not get Ids segment."), nil)
	}
	err = d.srv.Docker().ImageDelete(c, &b, metav1.DeleteOptions{})
	if err != nil {
		err = errors.Cause(err)
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"status": "ok"})
}
