package repositories

import (
	"post-system/lib/models"

	"gorm.io/gorm"
)

type tagsRepo struct {
	DB *gorm.DB
}

type TagsRepo interface {
	Insert(data models.Tag) error
	GetByLabel(label string) (*models.Tag, error)
}

func NewTagsRepo(db *gorm.DB) TagsRepo {
	return &tagsRepo{DB: db}
}

func (r *tagsRepo) Insert(data models.Tag) error {
    return r.DB.Create(&data).Error
}

func (r *tagsRepo) GetByLabel(label string) (*models.Tag, error) {
	var tag models.Tag
	err := r.DB.Where("label = ?", label).First(&tag).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &tag, err
}
