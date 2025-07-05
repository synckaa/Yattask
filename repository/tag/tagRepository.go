package tag

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TagRepository interface {
	Create(ctx context.Context, tx *gorm.DB, task model.Tag) (model.Tag, error)
	Update(ctx context.Context, tx *gorm.DB, task model.Tag) (model.Tag, error)
	Delete(ctx context.Context, tx *gorm.DB, task model.Tag) error
	GetById(ctx context.Context, tx *gorm.DB, taskId int) (model.Tag, error)
	GetAll(ctx context.Context, tx *gorm.DB) ([]model.Tag, error)
}
