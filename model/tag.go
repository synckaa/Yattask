package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"column:user_id"`
	Name   string `json:"name" gorm:"unique;column:name"`
	Tasks  []Task `json:"tasks" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:task_id"`
}

func (t *Tag) TableName() string {
	return "tags"
}
