package dto

import (
	"gorm.io/gorm"
)

type UserLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserServiceResponseGETWithTaskTag struct {
	gorm.Model
	Username string                 `json:"username"`
	Tasks    []TaskWithTagsResponse `json:"tasks"`
	Tags     []string               `json:"tags"`
}

type UserServiceResponse struct {
	gorm.Model
	Username string `json:"username"`
}
