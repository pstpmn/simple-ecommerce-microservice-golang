package pkg

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IGormORM interface {
	ConnectDB(dsn string, driver string) (*gorm.DB, error)
}

type GormORM struct{}

func (*GormORM) ConnectDB(dsn string, driver string) (*gorm.DB, error) {
	var gormDB *gorm.DB
	var err error

	switch driver {
	case "mysql":
		gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		return nil, errors.New("ErrUnsupportedDatabase")
	}

	return gormDB, err
}

func NewGormORM() IGormORM {
	return &GormORM{}
}
