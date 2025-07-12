package dto

import (
	"gorm.io/gorm"
)

type TaskCreateRequest struct {
	UserID      uint     `json:"user_id" form:"user_id" binding:"required"`
	Title       string   `json:"title" form:"title" binding:"required"`
	Deadline    string   `json:"deadline" form:"deadline" binding:"required"`
	Description string   `json:"description"  form:"description" binding:"required"`
	Status      bool     `json:"status" form:"status" binding:"required"`
	Tags        []string `json:"tags" form:"tags" binding:"required"`
}
type TaskUpdateRequest struct {
	ID          uint     `json:"id"  form:"id" binding:"required"`
	UserID      uint     `json:"user_id" form:"user_id" binding:"required"`
	Title       string   `json:"title" form:"title" binding:"required"`
	Deadline    string   `json:"deadline" form:"deadline" binding:"required"`
	Description string   `json:"description" form:"description" binding:"required"`
	Status      bool     `json:"status"`
	Tags        []string `json:"tags"`
}

type TaskServiceResponse struct {
	gorm.Model
	UserID      uint     `json:"user_id"`
	Title       string   `json:"title"`
	Deadline    string   `json:"deadline"`
	Description string   `json:"description"`
	Status      bool     `json:"status"`
	Tags        []string `json:"tags"`
}

type TaskWebCreateRequest struct {
	Title       string   `json:"title"`
	Deadline    string   `json:"deadline"`
	Description string   `json:"description"`
	Status      bool     `json:"status"`
	Tags        []string `json:"tags"`
}

type TaskWebUpdateRequest struct {
	Title       string   `json:"title"`
	Deadline    string   `json:"deadline"`
	Description string   `json:"description"`
	Status      bool     `json:"status"`
	Tags        []string `json:"tags"`
}

type TaskWithTagsResponse struct {
	gorm.Model
	Title       string   `json:"title"`
	Deadline    string   `json:"deadline"`
	Description string   `json:"description"`
	Status      bool     `json:"status"`
	Tags        []string `json:"tags"`
}
