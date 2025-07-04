package config

import (
	"Yattask/model"
	"gorm.io/gorm"
)

func SyncTables(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.Tag{})

}
