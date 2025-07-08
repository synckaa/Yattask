package helper

import (
	"Yattask/dto/taskDTO"
	"Yattask/model"
)

func ToTaskServiceResponse(task model.Task) taskDTO.TaskServiceResponse {
	return taskDTO.TaskServiceResponse{
		UserID:      task.UserID,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
	}
}
