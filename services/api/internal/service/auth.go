package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/DavelPurov777/microblog/services/api/internal/events"
	"github.com/DavelPurov777/microblog/services/api/internal/models"
)

type AuthService struct {
	repo                    AuthorizationRepo
	salt                    string
	userRegisteredPublisher events.UserRegisteredPublisher
}

func NewAuthService(repo AuthorizationRepo, salt string, userRegisteredPublisher events.UserRegisteredPublisher) *AuthService {
	return &AuthService{repo: repo, salt: salt, userRegisteredPublisher: userRegisteredPublisher}
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
	if err := s.userRegisteredPublisher.PublishUserRegistered(ev); err != nil {
		return 0, err
	}

	return id, nil
}

func generatePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
