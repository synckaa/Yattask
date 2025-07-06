package userDTO

import (
	"Yattask/model"
	"gorm.io/gorm"
)

type UserServiceResponseGETWithTaskTag struct {
	gorm.Model
	Username string       `json:"username" gorm:"column:username;unique"`
	Password string       `json:"password" gorm:"column:password"`
	Tasks    []model.Task `json:"tasks" gorm:"foreignKey:user_id;references:id"`
	Tags     []model.Tag  `json:"tags" gorm:"foreignKey:user_id;references:id"`
}
