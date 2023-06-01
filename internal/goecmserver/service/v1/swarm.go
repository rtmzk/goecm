package v1

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type SwarmSrv interface {
	Create(ctx context.Context, swarm *v1.SwarmNodeSpec, options metav1.CreateOptions) error
	Register(ctx context.Context, host *v1.HostInfo, options metav1.CreateOptions) error
	GetToken(ctx context.Context, role string, options metav1.GetOptions) (string, error)
	GenderToken(ctx context.Context, tokens *v1.SwarmJoin, options metav1.CreateOptions) error
}

type swarmService struct {
	store store.Factory
}

var _ SwarmSrv = (*swarmService)(nil)

func newSwarmSrv(srv *service) *swarmService {
	return &swarmService{
		store: srv.store,
	}
}

func (s *swarmService) Create(ctx context.Context, swarm *v1.SwarmNodeSpec, opts metav1.CreateOptions) error {
	return nil
}

func (s *swarmService) Register(ctx context.Context, info *v1.HostInfo, opts metav1.CreateOptions) error {
	return s.store.Swarm().Register(ctx, info, opts)
}

func (s *swarmService) GetToken(ctx context.Context, role string, opts metav1.GetOptions) (string, error) {
	return s.store.Swarm().GetToken(ctx, role, opts)
}

func (s *swarmService) GenderToken(ctx context.Context, tokens *v1.SwarmJoin, opts metav1.CreateOptions) error {
	return s.store.Swarm().GenderToken(ctx, tokens, opts)
}
