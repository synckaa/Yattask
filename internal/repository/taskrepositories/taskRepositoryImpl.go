package taskrepositories

import (
	"Yattask/internal/entities"
	"context"
	"errors"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (t TaskRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, task entities.Task) (entities.Task, error) {
	err := tx.WithContext(ctx).Create(&task).Error
	if err != nil {
		return entities.Task{}, err
	}
	var createdTask entities.Task
	err = tx.WithContext(ctx).Preload("Tags").Take(&createdTask, task.ID).Error
	if err != nil {
		return entities.Task{}, err
	}

	return createdTask, nil
}

func (t TaskRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, task entities.Task) (entities.Task, error) {
	err := tx.WithContext(ctx).Save(&task).Error
	if err != nil {
		return entities.Task{}, err
	}
	var updatedTask entities.Task
	err = tx.WithContext(ctx).Preload("Tags").Take(&updatedTask, task.ID).Error
	return updatedTask, nil
}

func (t TaskRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, taskId uint, userId uint) error {

	result := tx.WithContext(ctx).Where("id = ? AND user_id = ?", taskId, userId).Delete(&entities.Task{})
	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
