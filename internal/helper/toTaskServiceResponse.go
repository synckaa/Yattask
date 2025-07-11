package helper

import (
	"Yattask/internal/dto"
	"Yattask/internal/entities"
)

func ToTaskServiceResponse(task entities.Task) dto.TaskServiceResponse {
	var tags []string
	for _, tag := range task.Tags {
		tags = append(tags, tag.Name)
	}
	return dto.TaskServiceResponse{
		Model:       task.Model,
		UserID:      task.UserID,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
		Tags:        tags,
	}
}
