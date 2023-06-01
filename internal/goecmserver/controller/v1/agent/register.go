package agent

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
)

func (a *AgentController) Register(c *gin.Context) {
	var agent *v1.Agent
	var resp = new(v1.RegistryResponse)
	if err := c.ShouldBindJSON(&agent); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "Error bind request body. error: %s", err.Error()), nil)
		return
	}

	err := a.srv.Agent().Register(c, *agent)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp.Success = "ok"
	workerKey, err := a.srv.Swarm().GetToken(c, "worker", metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	resp.SwarmKey = workerKey

	core.WriteResponse(c, nil, resp)
}
