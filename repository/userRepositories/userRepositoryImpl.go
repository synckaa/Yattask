package userRepositories

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{}
}

func (u UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user model.User) (model.User, error) {
	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user model.User) (model.User, error) {
	err := tx.WithContext(ctx).Save(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) GetById(ctx context.Context, tx *gorm.DB, userId uint) (model.User, error) {
	var user model.User
	err := tx.WithContext(ctx).Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) GetByUsername(ctx context.Context, tx *gorm.DB, username string) (model.User, error) {
	var user model.User
	err := tx.WithContext(ctx).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) GetByIdWithTaskAndTag(ctx context.Context, tx *gorm.DB, userId uint) (model.User, error) {
	var user model.User
	err := tx.WithContext(ctx).Preload("Tasks.Tags").Preload("Tags").Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
