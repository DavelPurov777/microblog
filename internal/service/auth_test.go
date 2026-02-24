package service

import (
	"fmt"
	"testing"

	"github.com/DavelPurov777/microblog/internal/models"
)

type mockRepo struct{}

func TestService_generatePasswordHash(t *testing.T) {
	password := "123456"
	hash1 := generatePasswordHash(password)
	hash2 := generatePasswordHash(password)

	if hash1 != hash2 {
		t.Errorf("hashes should be equal")
	}
}

func TestService_CreateUser(t *testing.T) {
	repo := &mockRepo{}
	service := NewAuthService(repo)
	user := models.User{
		Password: "123456",
	}
	id, err := service.CreateUser(user)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if id != 1 {
		t.Fatalf("expected id 1 but got %v", id)
	}
}

func (m *mockRepo) CreateUser(user models.User) (int, error) {
	if user.Password == "123456" {
		return 0, fmt.Errorf("password not hashed")
	}

	return 1, nil
}
