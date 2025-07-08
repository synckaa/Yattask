package taskDTO

import "Yattask/model"

type TaskCreateUpdateRequest struct {
	UserID      uint        `json:"user_id" gorm:"column:user_id"`
	Title       string      `json:"title" gorm:"column:title"`
	Deadline    string      `json:"deadline" gorm:"column:deadline"`
	Description string      `json:"description" gorm:"column:description"`
	Status      bool        `json:"status" gorm:"column:status"`
	Tags        []model.Tag `json:"tags" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:task_id;references:id;joinReferences:tag_id"`
}
