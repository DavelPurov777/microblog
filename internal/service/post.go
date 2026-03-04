package service

import "github.com/DavelPurov777/microblog/internal/models"

type LikeEvent struct {
	PostID int
	UserID int
}

type LikeQueue interface {
	Publish(LikeEvent)
	Channel() <-chan LikeEvent
}

type PostListService struct {
	repo  PostsListRepo
	queue LikeQueue
}

func NewPostListService(repo PostsListRepo, q LikeQueue) *PostListService {
	return &PostListService{repo: repo, queue: q}
}

func (s *PostListService) Create(list models.Post) (int, error) {
	return s.repo.Create(list)
}

func (s *PostListService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}

func (s *PostListService) LikePost(postId, userId int) error {
	s.queue.Publish(LikeEvent{PostID: postId, UserID: userId})
	return nil
}

func (s *PostListService) processLike(ev LikeEvent) error {
	return s.repo.LikePost(ev.PostID, ev.UserID)
}

func (s *PostListService) StartLikeWorker(logger Logger) {
	go func() {
		for ev := range s.queue.Channel() {
			if err := s.processLike(ev); err != nil {
				logger.Error(err.Error())
			}
		}
	}()
}
