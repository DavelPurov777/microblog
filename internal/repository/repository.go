package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/DavelPurov777/microblog/internal/models"
)

const (
	usersTable = "users"
)

type Authorization interface {
	CreateUser(models.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}