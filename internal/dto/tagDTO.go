package dto

import (
	"Yattask/internal/entities"
	"gorm.io/gorm"
)

type TagServiceRequest struct {
	gorm.Model
	Name  string          `json:"name" gorm:"unique;column:name"`
	Tasks []entities.Task `json:"tasks" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:task_id"`
}

type TagServiceResponse struct {
	gorm.Model
	UserID uint            `json:"user_id" gorm:"column:user_id"`
	Name   string          `json:"name" gorm:"unique;column:name"`
	Tasks  []entities.Task `json:"tasks" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:task_id"`
}
