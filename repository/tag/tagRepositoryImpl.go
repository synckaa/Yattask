package tag

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
}

func (t TagRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, task model.Tag) (model.Tag, error) {
	err := tx.WithContext(ctx).Create(&task).Error
	if err != nil {
		return model.Tag{}, err
	}
	return task, nil
}

func (t TagRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, task model.Tag) (model.Tag, error) {
	err := tx.WithContext(ctx).Save(&task).Error
	if err != nil {
		return model.Tag{}, err
	}
	return task, nil
}

func (t TagRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, task model.Tag) error {
	err := tx.WithContext(ctx).Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (t TagRepositoryImpl) GetById(ctx context.Context, tx *gorm.DB, taskId int) (model.Tag, error) {
	var tag model.Tag
	err := tx.WithContext(ctx).Where("id = ?", taskId).Take(&tag).Error
	if err != nil {
		return model.Tag{}, err
	}
	return tag, nil
}

func (t TagRepositoryImpl) GetAll(ctx context.Context, tx *gorm.DB) ([]model.Tag, error) {
	var tags []model.Tag
	err := tx.WithContext(ctx).Find(&tags).Error
	if err != nil {
		return []model.Tag{}, err
	}
	return tags, nil
}
