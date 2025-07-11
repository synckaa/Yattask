package userrepositories

import (
	"Yattask/internal/entities"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error)
	Update(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error)
	GetById(ctx context.Context, tx *gorm.DB, userId uint) (entities.User, error)
	GetByUsername(ctx context.Context, tx *gorm.DB, username string) (entities.User, error)
	GetByIdWithTaskAndTag(ctx context.Context, tx *gorm.DB, userId uint) (entities.User, error)
}
