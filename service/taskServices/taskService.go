package taskServices

import (
	"Yattask/dto/taskDTO"
	"context"
)

type TaskService interface {
	Create(ctx context.Context, task taskDTO.TaskCreateUpdateRequest) (taskDTO.TaskServiceResponse, error)
	Update(ctx context.Context, task taskDTO.TaskCreateUpdateRequest) (taskDTO.TaskServiceResponse, error)
	Delete(ctx context.Context, taskID uint) error
	GetByID(ctx context.Context, taskID uint) (taskDTO.TaskServiceResponse, error)
}
