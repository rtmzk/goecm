package v1

import (
	"github.com/docker/docker/client"
	"go-ecm/internal/goecmserver/store"
)

type Service interface {
	Users() UserSrv
	Swarm() SwarmSrv
	Menu() MenuSrv
	SystemConfig() SystemConfigSrv
	Agent() AgentSrv
	Docker() DockerSrv
	SSHkey() SshSrv
}

type service struct {
	store store.Factory
	cli   *client.Client
}

var _ Service = (*service)(nil)

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func NewDockerService(store store.Factory, cli *client.Client) Service {
	return &service{
		store: store,
		cli:   cli,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Swarm() SwarmSrv {
	return newSwarmSrv(s)
}

func (s *service) Menu() MenuSrv {
	return newMenuSrv(s)
}

func (s *service) Docker() DockerSrv {
	return NewDockerSrv(s)
}

func (s *service) SSHkey() SshSrv {
	return newSshSrv(s)
}

func (s *service) SystemConfig() SystemConfigSrv {
	return newSystemConfigSrv(s)
}

func (s *service) Agent() AgentSrv {
	return newAgentSrv(s)
}
