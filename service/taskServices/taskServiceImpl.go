package taskServices

import (
	"Yattask/dto/taskDTO"
	"Yattask/helper"
	"Yattask/model"
	"Yattask/repository/tagRepositories"
	"Yattask/repository/taskRepositories"
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TaskServiceImpl struct {
	DB       *gorm.DB
	Repo     taskRepositories.TaskRepository
	tagRepo  tagRepositories.TagRepository
	Validate *validator.Validate
}

func NewTaskService(db *gorm.DB, repo taskRepositories.TaskRepository, tagRepo tagRepositories.TagRepository, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		DB:       db,
		Repo:     repo,
		tagRepo:  tagRepo,
		Validate: validate,
	}
}

func (t TaskServiceImpl) Create(ctx context.Context, task taskDTO.TaskCreateUpdateRequest) (taskDTO.TaskServiceResponse, error) {
	err := t.Validate.Struct(task)
	if err != nil {
		return taskDTO.TaskServiceResponse{}, err
	}
	taskReq := model.Task{
		UserID:      task.UserID,
		Title:       task.Title,
		Deadline:    task.Deadline,
		Description: task.Description,
		Status:      task.Status,
	}
	var taskResp model.Task
	errTX := t.DB.Transaction(func(tx *gorm.DB) error {
		for _, tagName := range task.Tags {
			tag, err := t.tagRepo.FindByName(ctx, tx, tagName.Name, task.UserID)
			if err != nil {
				return err
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
		return taskDTO.TaskServiceResponse{}, errTX
	}
	return helper.ToTaskServiceResponse(taskResp), nil
}

func (t TaskServiceImpl) Update(ctx context.Context, task taskDTO.TaskCreateUpdateRequest) (taskDTO.TaskServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) Delete(ctx context.Context, taskID uint) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskServiceImpl) GetByID(ctx context.Context, taskID uint) (taskDTO.TaskServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}
