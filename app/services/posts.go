package services

import (
	"errors"
	"post-system/app/repositories"
	"post-system/lib/models"
)

type postsService struct {
	postsRepo repositories.PostsRepo
	tagsRepo  repositories.TagsRepo
}

type PostsService interface {
	Insert(payload models.AddPost) error
	GetAll() ([]models.Post, error)
	GetById(id int) (*models.Post, error)
	Update(id int, payload models.UpdatePost) error
	Delete(id int) error
}

func NewPostsService(postsRepo repositories.PostsRepo, tagsRepo repositories.TagsRepo) PostsService {
	return &postsService{
		postsRepo: postsRepo,
		tagsRepo:  tagsRepo,
	}
}

func (s *postsService) Insert(payload models.AddPost) error {
	postData := models.Post{Title: payload.Title, Content: payload.Content}

	for _, label := range payload.Tags {
		tag, err := s.tagsRepo.GetByLabel(label)
		if err != nil {
			return err
		}
		if tag == nil {
			tag = &models.Tag{Label: label}
			err = s.tagsRepo.Insert(*tag)
			if err != nil {
				return err
			}
		}

		postData.Tags = append(postData.Tags, *tag)
	}

	return s.postsRepo.Insert(postData)
}

func (s *postsService) GetAll() ([]models.Post, error) {
	return s.postsRepo.GetAllPreloaded()
}

func (s *postsService) GetById(id int) (*models.Post, error) {
	return s.postsRepo.GetByIdPreloaded(id)
}

func (s *postsService) Update(id int, payload models.UpdatePost) error {
	postData, err := s.postsRepo.GetById(id)
	if err != nil {
		return err
	}

	postData.Title = payload.Title
	postData.Content = payload.Content

	var newTags []models.Tag
	for _, label := range payload.Tags {
		tag, err := s.tagsRepo.GetByLabel(label)
		if err != nil {
			return err
		}
		if tag == nil {
			tag = &models.Tag{Label: label}
			err = s.tagsRepo.Insert(*tag)
			if err != nil {
				return err
			}
		}

		newTags = append(newTags, *tag)
	}

	if err := s.postsRepo.ReplaceAssociation(*postData, newTags); err != nil {
		return err
	}

	if err := s.postsRepo.Update(*postData); err != nil {
		return err
	}

	return nil
}

func (s *postsService) Delete(id int) error {
	postData, err := s.postsRepo.GetById(id)
	if err != nil {
		return err
	}
	if postData != nil {
		return s.postsRepo.Delete(id)
	} else {
		return errors.New("post data not found")
	}
}
