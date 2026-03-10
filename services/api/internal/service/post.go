package service

import (
	"time"

	"github.com/DavelPurov777/microblog/services/api/internal/events"
	"github.com/DavelPurov777/microblog/services/api/internal/models"
)

type PostListService struct {
	repo      PostsListRepo
	publisher events.LikeEventPublisher
}

func NewPostListService(repo PostsListRepo, publisher events.LikeEventPublisher) *PostListService {
	return &PostListService{repo: repo, publisher: publisher}
}

func (s *PostListService) Create(list models.Post) (int, error) {
	return s.repo.Create(list)
}

func (s *PostListService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}

func (s *PostListService) LikePost(postId, userId int) error {
	ev := events.LikeEvent{
		PostID:    postId,
		UserID:    userId,
		CreatedAt: time.Now(),
	}

	return s.publisher.PublishLike(ev)
}
