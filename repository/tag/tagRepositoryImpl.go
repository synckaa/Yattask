package tag

import (
	"Yattask/model"
	"context"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
}

func (t TagRepositoryImpl) FindOrCreateByName(ctx context.Context, tx *gorm.DB, name string, userId uint) (model.Tag, error) {
	var tag model.Tag
	err := tx.WithContext(ctx).Where("name = ? AND user_id= ?", name, userId).First(&tag).Error
	if err != nil {
		tag = model.Tag{Name: name, UserID: userId}
		err = tx.WithContext(ctx).Create(&tag).Error
		if err != nil {
			return model.Tag{}, err
		}
	}
	return tag, nil

}
