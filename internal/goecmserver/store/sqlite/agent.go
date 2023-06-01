package sqlite

import (
	"context"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/pkg/code"
	metav1 "go-ecm/internal/pkg/meta/v1"
	"go-ecm/pkg/log"
	"gorm.io/gorm"
	"time"
)

type agent struct {
	db *gorm.DB
}

func newAgent(ds *datastore) *agent {
	return &agent{
		db: ds.db,
	}
}

func (a *agent) Register(ctx context.Context, agent v1.Agent) error {
	var count int64
	a.db.Table("agents").Where("address = ? and port = ?", agent.AgentAddr, agent.AgentPort).Count(&count)
	if count > 0 {
		return errors.WithCode(code.ErrAlreadyRegister, "节点已注册")
	}

	if err := a.db.Table("agents").Create(&agent).Error; err != nil {
		return errors.WithCode(code.ErrDatabase, "无法创建新节点记录")
	}
	return nil
}

func (a *agent) Heartbeat(ctx context.Context, agents v1.Agent) error {

	return nil
}

func (a *agent) GetAgents(ctx context.Context, options metav1.GetOptions) ([]v1.Agent, error) {
	t := time.Now()
	log.L(ctx).Debugf("get agent list from db. start time: %T", time.Now())
	var agents []v1.Agent
	err := a.db.Table("agents").Find(&agents).Error
	if err != nil {
		return nil, errors.WithCode(code.ErrDatabase, "无法获取agent列表")
	}
	log.L(ctx).Debugf("get agent list from db end. total cost time: %T", time.Since(t))
	return agents, nil
}

func (a *agent) UpdateAgents(ctx context.Context, options metav1.UpdateOptions, agent *v1.Agent) error {
	return a.db.Save(agent).Error
}
