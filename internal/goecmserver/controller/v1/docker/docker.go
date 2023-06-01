package docker

import (
	"github.com/docker/docker/client"
	srvv1 "go-ecm/internal/goecmserver/service/v1"
	"go-ecm/internal/goecmserver/store"
)

type DockerController struct {
	srv srvv1.Service
}

func NewDockerController(store store.Factory, dockerCliIns *client.Client) *DockerController {
	return &DockerController{
		srv: srvv1.NewDockerService(store, dockerCliIns),
	}
}
