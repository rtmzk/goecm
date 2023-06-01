package mysql

import (
	"gorm.io/gorm"
)

type docker struct {
	db *gorm.DB
}

func newDocker(ds *datastore) *docker {
	return &docker{
		db: ds.db,
	}
}
