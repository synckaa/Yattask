package taskRepositories

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(ctx context.Context, tx *gorm.DB, task model.Task) (model.Task, error)
	Update(ctx context.Context, tx *gorm.DB, task model.Task) (model.Task, error)
	Delete(ctx context.Context, tx *gorm.DB, task model.Task) error
	GetByIdWithTags(ctx context.Context, tx *gorm.DB, taskId int) (model.Task, error)
	GetAll(ctx context.Context, tx *gorm.DB) ([]model.Task, error)
}
