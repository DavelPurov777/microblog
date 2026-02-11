package service

import (
	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/DavelPurov777/microblog/internal/repository"
)

type PostListService struct {
	repo repository.PostsList
}

func NewPostListService(repo repository.PostsList) *PostListService {
	return &PostListService{repo: repo}
} 

func (s *PostListService) Create(list models.Post)  (int, error) {
	return s.repo.Create(list)
}