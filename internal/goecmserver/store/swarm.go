package store

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type SwarmStore interface {
	Create(ctx context.Context, swarm *v1.SwarmNodeSpec, opts metav1.CreateOptions) error
	Update(ctx context.Context, swarm *v1.SwarmNodeSpec, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, node string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, node string, opts metav1.GetOptions) (*v1.SwarmNodeSpec, error)
	Register(ctx context.Context, host *v1.HostInfo, opts metav1.CreateOptions) error
	NodeList(ctx context.Context, opts metav1.ListOptions) (*v1.NodeList, error)
	GetToken(ctx context.Context, role string, opts metav1.GetOptions) (string, error)
	GenderToken(ctx context.Context, tokens *v1.SwarmJoin, opts metav1.CreateOptions) error
}
