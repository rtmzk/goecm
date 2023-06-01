package agent

import (
	srvv1 "go-ecm/internal/goecmserver/service/v1"
	"go-ecm/internal/goecmserver/store"
)

type AgentController struct {
	srv srvv1.Service
}

func NewAgentController(s store.Factory) *AgentController {
	return &AgentController{
		srv: srvv1.NewService(s),
	}
}
