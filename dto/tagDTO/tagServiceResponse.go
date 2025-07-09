package tagDTO

import (
	"Yattask/model"
	"gorm.io/gorm"
)

type TagServiceResponse struct {
	gorm.Model
	UserID uint         `json:"user_id" gorm:"column:user_id"`
	Name   string       `json:"name" gorm:"unique;column:name"`
	Tasks  []model.Task `json:"tasks" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:task_id"`
}
