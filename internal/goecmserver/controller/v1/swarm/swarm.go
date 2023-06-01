package swarm

import (
	srvv1 "go-ecm/internal/goecmserver/service/v1"
	"go-ecm/internal/goecmserver/store"
)

type SwarmController struct {
	svc srvv1.Service
}

func NewSwarmController(store store.Factory) *SwarmController {
	return &SwarmController{
		svc: srvv1.NewService(store),
	}
}
