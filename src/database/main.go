package database

import (
	"github.com/williamroberttv/curriculum-gen-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database struct {
	Db *gorm.DB
	Dsn string
	DbType string
	AutoMigrateDb bool
}

func NewDb() *Database {
	return &Database{}
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	d.Db, err = gorm.Open(postgres.Open(d.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if d.AutoMigrateDb {
		err = d.Db.AutoMigrate(&models.User{})
		if err != nil {
			return nil, err
		}
	}

	DB = d.Db
	return d.Db, nil
}