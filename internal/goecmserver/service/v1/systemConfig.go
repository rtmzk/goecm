package v1

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type SystemConfigSrv interface {
	GetSystemConfig(c context.Context, options metav1.GetOptions) (*v1.SystemConfig, error)
	GetAgentToken(c context.Context, options metav1.GetOptions) (string, error)
}

type systemConfigService struct {
	store store.Factory
}

func newSystemConfigSrv(srv *service) *systemConfigService {
	return &systemConfigService{
		store: srv.store,
	}
}

var _ SystemConfigSrv = (*systemConfigService)(nil)

func (s *systemConfigService) GetSystemConfig(c context.Context, options metav1.GetOptions) (*v1.SystemConfig, error) {
	return s.store.SystemConfig().GetSystemConfig(c, options)
}

func (s *systemConfigService) GetAgentToken(c context.Context, options metav1.GetOptions) (string, error) {
	return s.store.SystemConfig().GetAgentToken(c, options)
}
