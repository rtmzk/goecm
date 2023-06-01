package db

import (
	"go-ecm/internal/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type SqliteOptions struct {
	DBPath                string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

func NewSqlite(opts *SqliteOptions) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(opts.DBPath), &gorm.Config{
		Logger: logger.New(opts.LogLevel),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
