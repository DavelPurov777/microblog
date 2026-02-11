package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/DavelPurov777/microblog/internal/models"
)

const (
	usersTable = "users"
	postsListsTable = "posts_lists"
)

type Authorization interface {
	CreateUser(models.User) (int, error)
}

type PostsList interface {
	Create(list models.Post)  (int, error)
	GetAll() ([]models.Post, error)
	LikePost(listId int) error
}

type Repository struct {
	Authorization
	PostsList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		PostsList: NewPostListPostgres(db), 
	}
}