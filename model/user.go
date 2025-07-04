package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;unique"`
	Password string `json:"password" gorm:"column:password"`
	Tasks    []Task `json:"tasks" gorm:"foreignKey:user_id;references:id"`
	Tags     []Tag  `json:"tags" gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
