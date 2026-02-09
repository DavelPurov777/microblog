package service

import (
	"github.com/DavelPurov777/microblog/internal/repository"
	"github.com/DavelPurov777/microblog/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}