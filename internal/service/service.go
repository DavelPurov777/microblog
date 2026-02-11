package service

import (
	"github.com/DavelPurov777/microblog/internal/repository"
	"github.com/DavelPurov777/microblog/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type PostsList interface {
	Create(list models.Post)  (int, error)
}

type Service struct {
	Authorization
	PostsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		PostsList: NewPostListService(repos.PostsList),
	}
}