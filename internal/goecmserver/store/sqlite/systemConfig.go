package sqlite

import (
	"context"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"gorm.io/gorm"
)

type systemConfig struct {
	db *gorm.DB
}

func newSystemConfig(ds *datastore) *systemConfig {
	return &systemConfig{
		db: ds.db,
	}
}

func (sc *systemConfig) GetSystemConfig(c context.Context, options metav1.GetOptions) (*v1.SystemConfig, error) {
	var systemConfig v1.SystemConfig

	err := sc.db.Find(&systemConfig).Error
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, "数据库错误,无法找到system_config表")
	}

	return &systemConfig, nil
}

func (sc *systemConfig) GetAgentToken(c context.Context, options metav1.GetOptions) (string, error) {
	systemConfig, err := sc.GetSystemConfig(c, options)
	if err != nil {
		return "", err
	}
	return systemConfig.Token, nil
}
