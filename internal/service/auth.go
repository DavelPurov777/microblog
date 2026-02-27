package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/DavelPurov777/microblog/internal/models"
)

type AuthService struct {
	repo AuthorizationRepo
	salt string
}

func NewAuthService(repo AuthorizationRepo, salt string) *AuthService {
	return &AuthService{repo: repo, salt: salt}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password, s.salt)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
