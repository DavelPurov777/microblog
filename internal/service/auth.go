package service
 
import (
	"fmt"
	"crypto/sha1"
	"github.com/DavelPurov777/microblog/internal/repository"
	"github.com/DavelPurov777/microblog/internal/models"
)

const (
	salt = "fasdfasdfasdf"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}