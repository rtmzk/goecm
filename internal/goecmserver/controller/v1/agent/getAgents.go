package agent

import (
	"github.com/gin-gonic/gin"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/core"
)

func (a *AgentController) GetAgents(c *gin.Context) {
	data, err := a.srv.Agent().GetAgents(c, metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, data)
}
