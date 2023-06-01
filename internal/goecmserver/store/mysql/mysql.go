package mysql

import (
	"fmt"
	"github.com/marmotedu/errors"
	v1 "go-ecm/internal/goecmserver/model/v1"
	"go-ecm/internal/goecmserver/store"
	metav1 "go-ecm/internal/pkg/meta/v1"
	genericoptions "go-ecm/internal/pkg/options"
	"go-ecm/pkg/db"
	"gorm.io/gorm"
	"sync"
	"time"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) SSHKey() store.SSHkeyStore {
	return newSshKey(ds)
}

func (ds *datastore) Swarm() store.SwarmStore {
	return newSwarm(ds)
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Menu() store.MenuStore {
	return newMenu(ds)
}

func (ds *datastore) Docker() store.DockerStore {
	return newDocker(ds)
}

func (ds *datastore) SystemConfig() store.SystemConfigStore {
	return newSystemConfig(ds)
}

func (ds *datastore) Agent() store.AgentStore {
	return newAgent(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store factory")
	}

	var err error
	var dbIns *gorm.DB
	if err != nil {
		return nil, err
	}

	once.Do(func() {
		options := &db.MysqlOptions{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
		}
		dbIns, err = db.New(options)
		initDB(dbIns)
		mysqlFactory = &datastore{db: dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store factory, mysqlFactory: %v,error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

func initDB(db *gorm.DB) {
	db.AutoMigrate(
		v1.HostInfo{},
		v1.SwarmJoin{},
		v1.Agent{},
		v1.MenuBaseSpec{},
		v1.User{},
		v1.Status{},
	)
	var count int64
	//var defaultPassword = util.BcryptHash("edoc2")
	db.Table("menu_base_specs").Count(&count)
	if count == 0 {
		menus := []v1.MenuBaseSpec{
			{
				Id: "1", Name: "node_manager", Parent: "0", Type: 2, Url: "/main/index", ZhName: "节点管理", Children: nil, Icon: "nodem",
			},
			{
				Id: "2", Name: "image_manage", Parent: "0", Type: 2, Url: "/main/images", ZhName: "镜像管理", Children: nil, Icon: "imagem",
			},
			{

				Id: "3", Name: "container_manager", Parent: "0", Type: 2, Url: "/main/containers", ZhName: "容器管理", Children: nil, Icon: "containerm",
			},
		}

		user := []v1.User{
			{
				ObjectMeta: metav1.ObjectMeta{
					ID:        1,
					Name:      "admin",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				UserName: "admin",
				Password: "edoc2",
				NickName: "系统用户",
			},
		}

		db.Table("menu_base_specs").Create(&menus)
		db.Table("users").Create(&user)
	}
}
