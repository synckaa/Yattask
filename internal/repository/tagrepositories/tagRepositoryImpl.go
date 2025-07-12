package tagrepositories

import (
	"Yattask/internal/entities"
	"context"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
}

func NewTagRepository() TagRepository {
	return &TagRepositoryImpl{}
}

func (t *TagRepositoryImpl) FindByName(ctx context.Context, tx *gorm.DB, name string, userId uint) (entities.Tag, error) {
	var tag entities.Tag
	err := tx.WithContext(ctx).Where("name = ? AND user_id= ?", name, userId).Take(&tag).Error
	if err != nil {
		tag = entities.Tag{
			Name:   name,
			UserID: userId,
		}
		tx.Create(&tag)
	}
	return tag, nil

}

func (t *TagRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, tag entities.Tag) (entities.Tag, error) {
	err := tx.WithContext(ctx).Create(&tag).Error
	if err != nil {
		return entities.Tag{}, err
	}
	return tag, nil
}

func (t *TagRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, userId uint) error {
	subQuery := tx.
		WithContext(ctx).
		Table("tasks_tags AS tt").
		Joins("JOIN tasks t ON t.id = tt.task_id AND t.deleted_at IS NULL").
		Select("DISTINCT tt.tag_id")

	var unusedTags []entities.Tag
	err := tx.
		WithContext(ctx).
		Where("user_id = ?", userId).
		Where("id NOT IN (?)", subQuery).
		Find(&unusedTags).Error

	if err != nil {
		return err
	}

	for _, tag := range unusedTags {
		if err := tx.WithContext(ctx).Delete(&tag).Error; err != nil {
			return err
		}
	}

	return nil
}
