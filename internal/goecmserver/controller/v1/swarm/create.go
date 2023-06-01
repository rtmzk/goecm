package swarm

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
	"go-ecm/pkg/log"
)

func (s *SwarmController) Create(c *gin.Context) {
	log.L(c).Debug("swarm create function called.")

	var r *v1.SwarmNodeSpec
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "客户端传参错误"), nil)
		return
	}
}
