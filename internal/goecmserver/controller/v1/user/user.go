package user

import (
	srvv1 "go-ecm/internal/goecmserver/service/v1"
	"go-ecm/internal/goecmserver/store"
)

type UserController struct {
	srv srvv1.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}
