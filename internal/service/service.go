package service

import (
	"github.com/DavelPurov777/microblog/internal/models"
)

type AuthorizationRepo interface {
	CreateUser(models.User) (int, error)
}

type PostsListRepo interface {
	Create(list models.Post) (int, error)
	GetAll() ([]models.Post, error)
	LikePost(listId int) error
}

type Repositories interface {
	AuthorizationRepo
	PostsListRepo
}

type Logger interface {
	Error(string)
	Info(string)
}

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type PostsList interface {
	Create(list models.Post) (int, error)
	GetAll() ([]models.Post, error)
	LikePost(listId int) error
	processLike(listId int) error
	StartLikeWorker(logger Logger)
}

type Service struct {
	Authorization
	PostsList
}

func NewService(repos Repositories, likeQueue LikeQueue, salt string) *Service {
	return &Service{
		Authorization: NewAuthService(repos, salt),
		PostsList:     NewPostListService(repos, likeQueue),
	}
}
