package sqlite

import (
	"context"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"gorm.io/gorm"
)

type swarm struct {
	db *gorm.DB
}

func newSwarm(ds *datastore) *swarm {
	return &swarm{ds.db}
}

func (s *swarm) Create(ctx context.Context, swarm *v1.SwarmNodeSpec, opts metav1.CreateOptions) error {

	return nil
}

func (s *swarm) Delete(ctx context.Context, node string, opts metav1.DeleteOptions) error {

	return nil
}

func (s *swarm) Get(ctx context.Context, node string, opts metav1.GetOptions) (*v1.SwarmNodeSpec, error) {

	return nil, nil
}

func (s *swarm) Update(ctx context.Context, swarm *v1.SwarmNodeSpec, options metav1.UpdateOptions) error {
	return nil
}

func (s *swarm) Register(ctx context.Context, host *v1.HostInfo, opts metav1.CreateOptions) error {
	if err := s.db.Create(&host).Error; err != nil {
		return errors.WithCode(code.ErrDatabase, "注册失败，数据库错误")
	}
	return nil
}

func (s *swarm) NodeList(ctx context.Context, opts metav1.ListOptions) (*v1.NodeList, error) {
	ret := &v1.NodeList{}

	d := s.db.Select("host_ip").Find(&ret.Items)

	return ret, d.Error
}

func (s *swarm) GetToken(ctx context.Context, role string, opts metav1.GetOptions) (string, error) {
	tokens := &v1.SwarmJoin{}

	d := s.db.First(tokens)

	switch role {
	case "manager":
		return tokens.MgrJoinToken, d.Error
	case "worker":
		return tokens.WrkJoinToken, d.Error
	}

	return "", nil
}

func (s *swarm) GenderToken(ctx context.Context, tokens *v1.SwarmJoin, opts metav1.CreateOptions) error {
	if err := s.db.Create(&tokens).Error; err != nil {
		return errors.WithCode(code.ErrDatabase, "生成失败,数据库插入错误")
	}
	return nil
}
