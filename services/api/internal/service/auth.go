package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/DavelPurov777/microblog/services/api/internal/events"
	"github.com/DavelPurov777/microblog/services/api/internal/models"
)

type AuthService struct {
	repo          AuthorizationRepo
	salt          string
	userPublisher events.UserEventPublisher
}

func NewAuthService(repo AuthorizationRepo, salt string, userPublisher events.UserEventPublisher) *AuthService {
	return &AuthService{repo: repo, salt: salt, userPublisher: userPublisher}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password, s.salt)
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	ev := events.UserRegisteredEvent(events.UserRegisteredEvent{
		Id:       id,
		Username: user.Username,
	})
	if err := s.userPublisher.PublishUserRegistered(ev); err != nil {
		return 0, err
	}

	return id, nil
}

func generatePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
