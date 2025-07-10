package taskDTO

import (
	"gorm.io/gorm"
)

type TaskServiceResponse struct {
	gorm.Model
	UserID      uint     `json:"user_id" gorm:"column:user_id"`
	Title       string   `json:"title" gorm:"column:title"`
	Deadline    string   `json:"deadline" gorm:"column:deadline"`
	Description string   `json:"description" gorm:"column:description"`
	Status      bool     `json:"status" gorm:"column:status"`
	Tags        []string `json:"tags" `
}
