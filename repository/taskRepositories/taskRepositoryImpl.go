package taskRepositories

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (t TaskRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, task model.Task) (model.Task, error) {
	err := tx.WithContext(ctx).Create(&task).Error
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (t TaskRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, task model.Task) (model.Task, error) {
	err := tx.WithContext(ctx).Save(&task).Error
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (t TaskRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, task model.Task) error {
	err := tx.WithContext(ctx).Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func (t TaskRepositoryImpl) GetByIdWithTags(ctx context.Context, tx *gorm.DB, taskId int) (model.Task, error) {
	var task model.Task
	err := tx.WithContext(ctx).Where("id = ?", taskId).Take(&task).Error
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (t TaskRepositoryImpl) GetAll(ctx context.Context, tx *gorm.DB) ([]model.Task, error) {
	var tasks []model.Task
	err := tx.WithContext(ctx).Find(&tasks).Error
	if err != nil {
		return []model.Task{}, err
	}
	return tasks, nil
}
