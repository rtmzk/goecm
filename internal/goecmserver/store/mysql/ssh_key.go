package mysql

import (
	"context"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/log"
	"gorm.io/gorm"
)

type sshKey struct {
	db *gorm.DB
}

func newSshKey(ds *datastore) *sshKey {
	return &sshKey{
		db: ds.db,
	}
}

func (s *sshKey) Get(ctx context.Context, info *v1.KeyInfo, options metav1.GetOptions) (int64, error) {
	var count int64
	if !errors.Is(s.db.Model(&info).Count(&count).Error, gorm.ErrRecordNotFound) && count == 1 {
		log.Warn("ssh key pair already generated")
		return count, nil
	}
	return 0, nil
}

func (s *sshKey) Create(ctx context.Context, info *v1.KeyInfo, options metav1.CreateOptions) error {
	err := s.db.Create(&info).Error
	if err != nil {
		return errors.WithCode(code.ErrDatabase, "创建失败，数据库错误")
	}

	return nil
}

func (s *sshKey) GetPrivate(ctx context.Context, options metav1.GetOptions) (*v1.KeyInfo, error) {
	info := &v1.KeyInfo{}
	err := s.db.Select("private_key").Find(&info).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrPrivateKeyNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}
	return info, nil
}
