package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/DavelPurov777/microblog/internal/models"
)

type PostListPostgres struct {
	db *sqlx.DB
}

func NewPostListPostgres(db *sqlx.DB) *PostListPostgres {
	return &PostListPostgres{db: db}
}

func (r *PostListPostgres) Create(post models.Post)  (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", postsListsTable)
	row := tx.QueryRow(createPostQuery, post.Title, post.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}