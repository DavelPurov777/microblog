package repository

import (
	"github.com/DavelPurov777/microblog/internal/service"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	postsListsTable = "posts_lists"
)

type Repository struct {
	service.AuthorizationRepo
	service.PostsListRepo
}

func NewRepository(db *sqlx.DB) service.Repositories {
	return &Repository{
		AuthorizationRepo: NewAuthPostgres(db),
		PostsListRepo:     NewPostListPostgres(db),
	}
}
