package service

import (
	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/DavelPurov777/microblog/internal/repository"
)

type LikeQueue interface {
	Publish(int)
	Channel() <-chan int
}

type PostListService struct {
	repo  repository.PostsList
	queue LikeQueue
}

func NewPostListService(repo repository.PostsList, q LikeQueue) *PostListService {
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

func (s *PostListService) processLike(listId int) error {
	return s.repo.LikePost(listId)
}

func (s *PostListService) StartLikeWorker(logger Logger) {
	go func() {
		for id := range s.queue.Channel() {
			if err := s.processLike(id); err != nil {
				logger.Error(err.Error())
			}
		}
	}()
}
