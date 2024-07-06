package db

import (
	"github.com/HuyVQ18411c/go-todo/src/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(config *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.GetDatabaseURL(true)), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	return db
}
