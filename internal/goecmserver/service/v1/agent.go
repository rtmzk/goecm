package v1

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type AgentSrv interface {
	Register(ctx context.Context, agent v1.Agent) error
	GetAgents(ctx context.Context, options metav1.GetOptions) ([]v1.Agent, error)
	UpdateAgents(ctx context.Context, options metav1.UpdateOptions, agent *v1.Agent) error
	UpdateInitStatus(status int) error
}

type agentService struct {
	store store.Factory
}

func newAgentSrv(s *service) *agentService {
	return &agentService{
		store: s.store,
	}
}

func (asrv *agentService) Register(ctx context.Context, agent v1.Agent) error {
	return asrv.store.Agent().Register(ctx, agent)
}

func (asrv *agentService) GetAgents(ctx context.Context, options metav1.GetOptions) ([]v1.Agent, error) {
	return asrv.store.Agent().GetAgents(ctx, options)
}

func (asrv *agentService) UpdateAgents(ctx context.Context, options metav1.UpdateOptions, agent *v1.Agent) error {
	return asrv.store.Agent().UpdateAgents(ctx, options, agent)
}

func (asrv *agentService) UpdateInitStatus(status int) error {
	return asrv.store.Agent().UpdateInitStatus(status)
}
