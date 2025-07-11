package taskrepositories

import (
	"Yattask/internal/entities"
	"context"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(ctx context.Context, tx *gorm.DB, task entities.Task) (entities.Task, error)
	Update(ctx context.Context, tx *gorm.DB, task entities.Task) (entities.Task, error)
	Delete(ctx context.Context, tx *gorm.DB, taskId uint, userId uint) error
}
