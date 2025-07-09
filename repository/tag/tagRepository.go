package tag

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TagRepository interface {
	FindOrCreateByName(ctx context.Context, tx *gorm.DB, name string, userId string) (model.Tag, error)
}
