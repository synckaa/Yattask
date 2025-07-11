package taskservices

import (
	"Yattask/internal/dto"
	"context"
)

type TaskService interface {
	Create(ctx context.Context, task dto.TaskCreateRequest) (dto.TaskServiceResponse, error)
	Update(ctx context.Context, task dto.TaskUpdateRequest) (dto.TaskServiceResponse, error)
	Delete(ctx context.Context, taskID uint, userId uint) error
}
