package service

import (
	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/DavelPurov777/microblog/internal/queue"

	"github.com/DavelPurov777/microblog/internal/repository"
)

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

func NewService(repos *repository.Repository, likeQueue *queue.LikeQueue) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		PostsList:     NewPostListService(repos.PostsList, likeQueue),
	}
}
