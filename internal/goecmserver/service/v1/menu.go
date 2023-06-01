package v1

import (
	"context"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type MenuSrv interface {
	Get(ctx context.Context, opts metav1.ListOptions) ([]v1.MenuBaseSpec, error)
}

type menuService struct {
	store store.Factory
}

var _ MenuSrv = (*menuService)(nil)

func newMenuSrv(srv *service) *menuService {
	return &menuService{
		store: srv.store,
	}
}

func (m *menuService) Get(ctx context.Context, opts metav1.ListOptions) ([]v1.MenuBaseSpec, error) {
	menus, err := m.store.Menu().Get(ctx, opts)
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return menus, nil
}
