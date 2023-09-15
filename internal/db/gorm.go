package db

import (
	"github.com/glebarez/sqlite"
	"github.com/xapier14/todo/internal/models"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func Open() error {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	if db == nil {
		panic("gorm.DB is nil")
	}
	return db
}

func AutoMigrate() error {
	return db.AutoMigrate(&models.Todo{}, &models.UserCredential{}, &models.UserSession{})
}