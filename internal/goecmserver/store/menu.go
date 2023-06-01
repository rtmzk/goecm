package store

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type MenuStore interface {
	Get(ctx context.Context, opts metav1.ListOptions) ([]v1.MenuBaseSpec, error)
}
