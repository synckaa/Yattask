package tagrepositories

import (
	"Yattask/internal/entities"
	"context"
	"gorm.io/gorm"
)

type TagRepository interface {
	FindByName(ctx context.Context, tx *gorm.DB, name string, userId uint) (entities.Tag, error)
	Create(ctx context.Context, tx *gorm.DB, tag entities.Tag) (entities.Tag, error)
	Delete(ctx context.Context, tx *gorm.DB, userId uint) error
	Update(ctx context.Context, tx *gorm.DB, task entities.Task, tags []entities.Tag) error
}
