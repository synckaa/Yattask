package tagServices

import (
	"Yattask/dto/tagDTO"
	"Yattask/repository/tag"
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TagServiceImpl struct {
	DB            *gorm.DB
	tagRepository tag.TagRepository
	Validate      *validator.Validate
}

func NewTagServiceImpl(db *gorm.DB, tagRepository tag.TagRepository, validate *validator.Validate) TagService {
	return &TagServiceImpl{
		DB:            db,
		tagRepository: tagRepository,
		Validate:      validate,
	}
}

func (t TagServiceImpl) FindOrCreate(ctx context.Context, task tagDTO.TagServiceRequest) (tagDTO.TagServiceResponse, error) {
	err := t.Validate.Struct(task)
	if err != nil {
		return tagDTO.TagServiceResponse{}, err
	}
	t.tagRepository.FindOrCreateByName()
}
