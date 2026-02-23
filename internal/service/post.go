package service

import (
	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/DavelPurov777/microblog/internal/queue"
	"github.com/DavelPurov777/microblog/internal/repository"
)

type PostListService struct {
	repo  repository.PostsList
	queue *queue.LikeQueue
}

func NewPostListService(repo repository.PostsList, q *queue.LikeQueue) *PostListService {
	return &PostListService{
		repo:  repo,
		queue: q,
	}
}

func (s *PostListService) Create(list models.Post) (int, error) {
	return s.repo.Create(list)
}

func (s *PostListService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}

func (s *PostListService) LikePost(listId int) error {
	s.queue.Publish(listId)
	return nil
}

func (s *PostListService) ProcessLike(listId int) error {
	return s.repo.LikePost(listId)
}
