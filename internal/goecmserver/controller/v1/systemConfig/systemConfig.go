package systemConfig

import (
	srvv1 "go-ecm/internal/goecmserver/service/v1"
	"go-ecm/internal/goecmserver/store"
)

type SystemConfigController struct {
	srv srvv1.Service
}

func NewSystemConfigController(store store.Factory) *SystemConfigController {
	return &SystemConfigController{
		srv: srvv1.NewService(store),
	}
}
