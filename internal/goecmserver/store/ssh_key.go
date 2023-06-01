package store

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type SSHkeyStore interface {
	Get(ctx context.Context, info *v1.KeyInfo, options metav1.GetOptions) (int64, error)
	Create(ctx context.Context, info *v1.KeyInfo, options metav1.CreateOptions) error
	GetPrivate(ctx context.Context, options metav1.GetOptions) (*v1.KeyInfo, error)
}
