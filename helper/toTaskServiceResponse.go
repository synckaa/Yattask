package helper

import (
	"Yattask/dto/taskDTO"
	"Yattask/model"
)

func ToTaskServiceResponse(task model.Task) taskDTO.TaskServiceResponse {
	var tags []string
	for _, tag := range task.Tags {
		tags = append(tags, tag.Name)
	}
	return taskDTO.TaskServiceResponse{
		Model:       task.Model,
		UserID:      task.UserID,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
		Tags:        tags,
	}
}
