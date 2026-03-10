package service

import "github.com/DavelPurov777/microblog/services/engagement/internal/repository"

type StatsService struct {
	repo repository.PostLikesRepository
}

func NewStatsService(r repository.PostLikesRepository) *StatsService {
	return &StatsService{repo: r}
}

func (s *StatsService) GetPostLikes(postID int) (int64, error) {
	return s.repo.GetLikes(postID)
}
