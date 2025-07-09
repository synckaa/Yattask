package tagDTO

import (
	"Yattask/model"
	"gorm.io/gorm"
)

type TagServiceRequest struct {
	gorm.Model
	Name  string       `json:"name" gorm:"unique;column:name"`
	Tasks []model.Task `json:"tasks" gorm:"many2many:tasks_tags;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:task_id"`
}
