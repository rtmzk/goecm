package store

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type AgentStore interface {
	Register(ctx context.Context, agent v1.Agent) error
	Heartbeat(ctx context.Context, agent v1.Agent) error
	GetAgents(ctx context.Context, options metav1.GetOptions) ([]v1.Agent, error)
	UpdateAgents(ctx context.Context, options metav1.UpdateOptions, agent *v1.Agent) error
	UpdateInitStatus(status int) error
}
