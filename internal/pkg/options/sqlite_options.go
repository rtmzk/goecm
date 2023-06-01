package options

import (
	"github.com/spf13/pflag"
	"go-ecm/utils"
	"time"
)

type SQLiteOptions struct {
	DBPath                string        `json:"path,omitempty" mapstructure:"path"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	LogLevel              int           `json:"log-level"                          mapstructure:"log-level"`
}

func NewSqliteOptions() *SQLiteOptions {
	return &SQLiteOptions{
		DBPath:                utils.UserHome() + "/.goecm.db",
		MaxIdleConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
		MaxOpenConnections:    100,
		LogLevel:              1,
	}
}

func (o *SQLiteOptions) Validate() []error {
	errs := []error{}

	return errs
}

func (o *SQLiteOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.DBPath, "sqlite.dbpath", o.DBPath, ""+
		"Path of database store.")

	fs.IntVar(&o.MaxOpenConnections, "sqlite.max-idle-connections", o.MaxIdleConnections, ""+
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&o.MaxOpenConnections, "sqlite.max-open-connections", o.MaxOpenConnections, ""+
		"Maximum opened connections allowed to connect to mysql.")

	fs.DurationVar(&o.MaxConnectionLifeTime, "sqlite.max-connection-left-time", o.MaxConnectionLifeTime, ""+
		"Max connection life time allowed to connect to mysql.")

	fs.IntVar(&o.LogLevel, "sqlite.log-mode", o.LogLevel, ""+
		"Specify gorm log level.")
}
