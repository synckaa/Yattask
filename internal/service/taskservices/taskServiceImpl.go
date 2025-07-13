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

func (t *TaskServiceImpl) Create(ctx context.Context, task dto.TaskCreateRequest) (dto.TaskServiceResponse, error) {
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

func (t *TaskServiceImpl) Update(ctx context.Context, task dto.TaskUpdateRequest) (dto.TaskServiceResponse, error) {
	err := t.Validate.Struct(task)
	if err != nil {
		return dto.TaskServiceResponse{}, err
	}
	var taskResp entities.Task
	errTX := t.DB.Transaction(func(tx *gorm.DB) error {
		existingTask, err := t.Repo.GetById(ctx, tx, task.ID, task.UserID)
		if err != nil {
			return err
		}
		existingTask.ID = task.ID
		existingTask.Title = task.Title
		existingTask.Deadline = task.Deadline
		existingTask.Description = task.Description
		existingTask.Status = task.Status

		_, err = t.Repo.Update(ctx, tx, existingTask)
		if err != nil {
			return err
		}
		var tags []entities.Tag
		for _, tagName := range task.Tags {
			tag, err := t.tagRepo.FindByName(ctx, tx, tagName, task.UserID)
			if err != nil {
				return err
			}
			tags = append(tags, tag)
		}
		err = t.tagRepo.Update(ctx, tx, existingTask, tags)
		if err != nil {
			return err
		}

		taskResp, err = t.Repo.GetByIdWithTags(ctx, tx, existingTask.ID, existingTask.UserID)
		if err != nil {
			return err
		}

		err = t.tagRepo.Delete(ctx, tx, taskResp.UserID)
		return nil
	})
	if errTX != nil {
		return dto.TaskServiceResponse{}, errTX
	}
	return helper.ToTaskServiceResponse(taskResp), nil

}

func (t *TaskServiceImpl) Delete(ctx context.Context, taskID uint, userId uint) error {
	errTx := t.DB.Transaction(func(tx *gorm.DB) error {
		err := t.Repo.Delete(ctx, tx, taskID, userId)
		if err != nil {
			return err
		}
		err = t.tagRepo.Delete(ctx, tx, userId)
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
