package userDTO

import "gorm.io/gorm"

type UserServiceResponse struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required"`
}
