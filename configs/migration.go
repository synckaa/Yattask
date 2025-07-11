package configs

import (
	"Yattask/internal/entities"
	"gorm.io/gorm"
)

func SyncTables(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{},
		&entities.Task{},
		&entities.Tag{})

}
