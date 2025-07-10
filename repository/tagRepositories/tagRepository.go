package tagRepositories

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TagRepository interface {
	FindByName(ctx context.Context, tx *gorm.DB, name string, userId uint) (model.Tag, error)
	Create(ctx context.Context, tx *gorm.DB, tag model.Tag) (model.Tag, error)
}
