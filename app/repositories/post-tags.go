package repositories

import (
	"post-system/lib/models"

	"gorm.io/gorm"
)

type postTagsRepo struct {
	DB *gorm.DB
}

type PostTagsRepo interface {
    DeleteTagsByPostID(id int) error
}

func NewPostTagsRepo(db *gorm.DB) PostTagsRepo {
    return &postTagsRepo{DB: db}
}

func (r *postTagsRepo) DeleteTagsByPostID(id int) error {
    return r.DB.Where("post_id = ?", id).Delete(&models.PostTag{}).Error
}