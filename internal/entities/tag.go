package entities

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"column:name"`
	Tasks  []Task `json:"tasks" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:task_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *Tag) TableName() string {
	return "tags"
}
