package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	"go-ecm/pkg/core"
)

func (a *AgentController) SetInitStatus(c *gin.Context) {
	var status *v1.Status
	err := c.ShouldBindJSON(&status)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "Bad Request"), nil)
		return
	}

	err = a.srv.Agent().UpdateInitStatus(status.Status)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "Update init status failed. Err: %s", err.Error()), nil)
		return
	}
	core.WriteResponse(c, nil, map[string]string{"success": "ok"})
}
