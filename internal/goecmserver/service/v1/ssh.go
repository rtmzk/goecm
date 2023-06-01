package v1

import (
	"context"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	metav1 "go-ecm/internal/pkg/meta/v1"
)

type SshSrv interface {
	Get(ctx context.Context, info *v1.KeyInfo, options metav1.GetOptions) (int64, error)
	Create(ctx context.Context, info *v1.KeyInfo, options metav1.CreateOptions) error
}

type sshService struct {
	store store.Factory
}

func newSshSrv(s *service) *sshService {
	return &sshService{
		store: s.store,
	}
}

var _ SshSrv = (*sshService)(nil)

func (s *sshService) Get(ctx context.Context, info *v1.KeyInfo, options metav1.GetOptions) (int64, error) {
	return s.store.SSHKey().Get(ctx, info, options)
}

func (s *sshService) Create(ctx context.Context, info *v1.KeyInfo, options metav1.CreateOptions) error {
	return s.store.SSHKey().Create(ctx, info, options)
}
