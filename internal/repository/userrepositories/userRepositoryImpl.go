package userrepositories

import (
	"Yattask/internal/entities"
	"context"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{}
}

func (u UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error) {
	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error) {
	err := tx.WithContext(ctx).Save(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) GetById(ctx context.Context, tx *gorm.DB, userId uint) (entities.User, error) {
	var user entities.User
	err := tx.WithContext(ctx).Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) GetByUsername(ctx context.Context, tx *gorm.DB, username string) (entities.User, error) {
	var user entities.User
	err := tx.WithContext(ctx).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) GetByIdWithTaskAndTag(ctx context.Context, tx *gorm.DB, userId uint) (entities.User, error) {
	var user entities.User
	err := tx.WithContext(ctx).Preload("Tasks.Tags", "tags.user_id = ? ", userId).
		Preload("Tags", "tags.user_id = ?", userId).Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
