package tagRepositories

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
}

func NewTagRepository() TagRepository {
	return &TagRepositoryImpl{}
}

func (t TagRepositoryImpl) FindByName(ctx context.Context, tx *gorm.DB, name string, userId uint) (model.Tag, error) {
	var tag model.Tag
	err := tx.WithContext(ctx).Where("name = ? AND user_id= ?", name, userId).Take(&tag).Error
	if err != nil {
		tag = model.Tag{
			Name:   name,
			UserID: userId,
		}
		tx.Create(&tag)
	}
	return tag, nil

}

func (t TagRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, tag model.Tag) (model.Tag, error) {
	var tagReq model.Tag
	tx.WithContext(ctx).Create(&tagReq)
	return tagReq, nil
}
