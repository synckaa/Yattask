package user

import "gorm.io/gorm"

type UserServiceResponse struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
