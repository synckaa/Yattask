package dto

import (
	"gorm.io/gorm"
)

type TaskCreateRequest struct {
	UserID      uint     `json:"user_id" gorm:"column:user_id"`
	Title       string   `json:"title" gorm:"column:title"`
	Deadline    string   `json:"deadline" gorm:"column:deadline"`
	Description string   `json:"description" gorm:"column:description"`
	Status      bool     `json:"status" gorm:"column:status"`
	Tags        []string `json:"tags"`
}
type TaskUpdateRequest struct {
	ID          uint     `json:"id" gorm:"column:id"`
	UserID      uint     `json:"user_id" gorm:"column:user_id"`
	Title       string   `json:"title" gorm:"column:title"`
	Deadline    string   `json:"deadline" gorm:"column:deadline"`
	Description string   `json:"description" gorm:"column:description"`
	Status      bool     `json:"status" gorm:"column:status"`
	Tags        []string `json:"tags"`
}

type TaskServiceResponse struct {
	gorm.Model
	UserID      uint     `json:"user_id" gorm:"column:user_id"`
	Title       string   `json:"title" gorm:"column:title"`
	Deadline    string   `json:"deadline" gorm:"column:deadline"`
	Description string   `json:"description" gorm:"column:description"`
	Status      bool     `json:"status" gorm:"column:status"`
	Tags        []string `json:"tags" `
}

type TaskWebCreateRequest struct {
	Title       string   `json:"title" gorm:"column:title"`
	Deadline    string   `json:"deadline" gorm:"column:deadline"`
	Description string   `json:"description" gorm:"column:description"`
	Status      bool     `json:"status" gorm:"column:status"`
	Tags        []string `json:"tags"`
}

type TaskWebUpdateRequest struct {
	Title       string   `json:"title" gorm:"column:title"`
	Deadline    string   `json:"deadline" gorm:"column:deadline"`
	Description string   `json:"description" gorm:"column:description"`
	Status      bool     `json:"status" gorm:"column:status"`
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
