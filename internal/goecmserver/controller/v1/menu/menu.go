package menu

import (
	srvv1 "go-ecm/internal/goecmserver/service/v1"
	"go-ecm/internal/goecmserver/store"
)

type MenuController struct {
	srv srvv1.Service
}

func NewMenuController(store store.Factory) *MenuController {
	return &MenuController{
		srv: srvv1.NewService(store),
	}
}
