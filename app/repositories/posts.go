package repositories

import (
	"post-system/lib/models"

	"gorm.io/gorm"
)

type postsRepo struct {
	DB *gorm.DB
}

type PostsRepo interface {
	Insert(data models.Post) error
	GetAll() ([]models.Post, error)
	GetById(id int) (*models.Post, error)
	GetAllPreloaded() ([]models.Post, error)
	GetByIdPreloaded(id int) (*models.Post, error)
	Update(data models.Post) error
	ReplaceAssociation(postData models.Post, newTags []models.Tag) error
	Delete(id int) error
}

func NewPostsRepo(db *gorm.DB) PostsRepo {
    return &postsRepo{DB: db}
}

func (r *postsRepo) Insert(data models.Post) error {
	return r.DB.Create(&data).Error
}

func (r *postsRepo) GetAll() ([]models.Post, error) {
    var posts []models.Post
    err := r.DB.Find(&posts).Error
    return posts, err
}

func (r *postsRepo) GetById(id int) (*models.Post, error) {
	var post models.Post
    err := r.DB.First(&post, id).Error
	if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    return &post, err
}

func (r *postsRepo) GetAllPreloaded() ([]models.Post, error) {
    var posts []models.Post
    err := r.DB.Preload("Tags").Find(&posts).Error
    return posts, err
}

func (r *postsRepo) GetByIdPreloaded(id int) (*models.Post, error) {
	var post models.Post
    err := r.DB.Preload("Tags").First(&post, id).Error
	if err == gorm.ErrRecordNotFound {
        return nil, nil
    }
    return &post, err
}

func (r *postsRepo) Update(data models.Post) error {
	return r.DB.Updates(&data).Error
}

func (r *postsRepo) Delete(id int) error {
    return r.DB.Delete(&models.Post{}, id).Error
}

func (r *postsRepo) ReplaceAssociation(postData models.Post, newTags []models.Tag) error {
	return r.DB.Model(&postData).Association("tags").Replace(newTags)
}