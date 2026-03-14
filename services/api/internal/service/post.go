package service

import (
	"time"

	"github.com/DavelPurov777/microblog/services/api/internal/events"
	"github.com/DavelPurov777/microblog/services/api/internal/models"
)

type PostListService struct {
	repo      PostsListRepo
	publisher events.EventPublisher
}

func NewPostListService(
	repo PostsListRepo,
	publisher events.EventPublisher,
) *PostListService {
	return &PostListService{repo: repo, publisher: publisher}
}

func (s *PostListService) Create(list models.Post) (int, error) {
	id, err := s.repo.Create(list)
	if err != nil {
		return 0, err
	}

	ev := events.PostCreatedEvent(events.PostCreatedEvent{
		Id:        id,
		UserId:    list.UserId,
		Title:     list.Title,
		CreatedAt: list.CreatedAt,
	})
	_ = s.publisher.PublishPostCreated(ev)

	return id, nil
}

func (s *PostListService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}

func (s *PostListService) LikePost(PostID, UserID int) error {
	ev := events.PostLikedEvent{
		PostID:    PostID,
		UserID:    UserID,
		CreatedAt: time.Now(),
	}

	return s.publisher.PublishPostLiked(ev)
}
