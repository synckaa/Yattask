package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"column:user_id"`
	Title       string `json:"title" gorm:"column:title"`
	Deadline    string `json:"deadline" gorm:"column:deadline"`
	Description string `json:"description" gorm:"column:description"`
	Status      bool   `json:"status" gorm:"column:status"`
	Tags        []Tag  `json:"tags" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:task_id;references:id;joinReferences:tag_id"`
}

func (t *Task) TableName() string {
	return "tasks"
}
