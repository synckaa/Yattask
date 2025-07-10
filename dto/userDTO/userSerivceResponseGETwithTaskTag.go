package userDTO

import (
	"Yattask/dto/taskDTO"
	"gorm.io/gorm"
)

type UserServiceResponseGETWithTaskTag struct {
	gorm.Model
	Username string                         `json:"username" gorm:"column:username;unique"`
	Tasks    []taskDTO.TaskWithTagsResponse `json:"tasks"`
	Tags     []string                       `json:"tags"`
}
