package taskservices

import (
	"Yattask/internal/dto"
	"Yattask/internal/entities"
	"Yattask/internal/helper"
	"Yattask/internal/repository/tagrepositories"
	"Yattask/internal/repository/taskrepositories"
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TaskServiceImpl struct {
	DB       *gorm.DB
	Repo     taskrepositories.TaskRepository
	tagRepo  tagrepositories.TagRepository
	Validate *validator.Validate
}

func NewTaskService(db *gorm.DB, repo taskrepositories.TaskRepository, tagRepo tagrepositories.TagRepository, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		DB:       db,
		Repo:     repo,
		tagRepo:  tagRepo,
		Validate: validate,
	}
}

func (t TaskServiceImpl) Create(ctx context.Context, task dto.TaskCreateRequest) (dto.TaskServiceResponse, error) {
	err := t.Validate.Struct(task)
	if err != nil {
		return dto.TaskServiceResponse{}, err
	}
	taskReq := entities.Task{
		UserID:      task.UserID,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
	}
	var taskResp entities.Task
	errTX := t.DB.Transaction(func(tx *gorm.DB) error {
		for _, tagName := range task.Tags {
			tag, err := t.tagRepo.FindByName(ctx, tx, tagName, task.UserID)
			if err != nil {
				tag = entities.Tag{
					Name:   tagName,
					UserID: task.UserID,
				}
				tag, err = t.tagRepo.Create(ctx, tx, tag)
				if err != nil {
					return err
				}
			}
			taskReq.Tags = append(taskReq.Tags, tag)

		}
		createdTask, err := t.Repo.Create(ctx, tx, taskReq)
		if err != nil {
			return err
		}
		taskResp = createdTask
		return nil
	})
	if errTX != nil {
		return dto.TaskServiceResponse{}, errTX
	}
	return helper.ToTaskServiceResponse(taskResp), nil
}

func (t TaskServiceImpl) Update(ctx context.Context, task dto.TaskUpdateRequest) (dto.TaskServiceResponse, error) {
	err := t.Validate.Struct(task)
	if err != nil {
		return dto.TaskServiceResponse{}, err
	}
	taskReq := entities.Task{
		UserID:      task.UserID,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
	}
	taskReq.ID = task.ID
	var taskResp entities.Task
	errTX := t.DB.Transaction(func(tx *gorm.DB) error {
		for _, tagName := range task.Tags {
			tag, err := t.tagRepo.FindByName(ctx, tx, tagName, task.UserID)
			if err != nil {
				return err
			}
			taskReq.Tags = append(taskReq.Tags, tag)
		}
		updatedTask, err := t.Repo.Update(ctx, tx, taskReq)
		if err != nil {
			return err
		}
		taskResp = updatedTask
		return nil
	})
	if errTX != nil {
		return dto.TaskServiceResponse{}, errTX
	}
	return helper.ToTaskServiceResponse(taskResp), nil

}

func (t TaskServiceImpl) Delete(ctx context.Context, taskID uint, userId uint) error {
	errTx := t.DB.Transaction(func(tx *gorm.DB) error {
		err := t.Repo.Delete(ctx, tx, taskID, userId)
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		return errTx
	}
	return nil
}
