package tagServices

import (
	"Yattask/dto/tagDTO"
	"Yattask/dto/taskDTO"
	"Yattask/model"
	"Yattask/repository/tagRepositories"
	"Yattask/repository/taskRepositories"
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TagServiceImpl struct {
	DB             *gorm.DB
	tagRepository  tagRepositories.TagRepository
	taskRepository taskRepositories.TaskRepository
	Validate       *validator.Validate
}

func NewTagServiceImpl(db *gorm.DB, tagRepository tagRepositories.TagRepository, taskRepository taskRepositories.TaskRepository, validate *validator.Validate) TagService {
	return &TagServiceImpl{
		DB:             db,
		tagRepository:  tagRepository,
		taskRepository: taskRepository,
		Validate:       validate,
	}
}

func (t TagServiceImpl) FindOrCreate(ctx context.Context, userId uint, input taskDTO.TaskCreateUpdateRequest) (tagDTO.TagServiceResponse, error) {
	err := t.Validate.Struct(input)
	if err != nil {
		return tagDTO.TagServiceResponse{}, err
	}

	task := model.Task{
		UserID:      userId,
		Title:       input.Title,
		Deadline:    input.Deadline,
		Description: input.Description,
		Status:      input.Status,
	}
	errTx := t.DB.Transaction(func(tx *gorm.DB) error {
		var tag model.Tag
		for _, tagName := range input.Tags {
			tag, err = t.tagRepository.FindByName(ctx, tx, tagName.Name, userId)
			if err != nil {
				return err
			}
			task.Tags = append(task.Tags, tag)
		}
		_, err = t.taskRepository.Create(ctx, tx, task)
		if err != nil {
			return err
		}
		return nil
	})
	if errTx != nil {
		return tagDTO.TagServiceResponse{}, errTx
	}

	return
}
