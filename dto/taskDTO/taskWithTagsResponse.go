package taskDTO

import "gorm.io/gorm"

type TaskWithTagsResponse struct {
	gorm.Model
	Title       string   `json:"title"`
	Deadline    string   `json:"deadline"`
	Description string   `json:"description"`
	Status      bool     `json:"status"`
	Tags        []string `json:"tags"`
}
