package user

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user model.User) (model.User, error)
	Update(ctx context.Context, tx *gorm.DB, user model.User) (model.User, error)
	GetById(ctx context.Context, tx *gorm.DB, userId int) (model.User, error)
	GetByUsername(ctx context.Context, tx *gorm.DB, username string) (model.User, error)
}
