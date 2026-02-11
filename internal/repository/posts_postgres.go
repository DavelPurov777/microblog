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
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, description, likes) VALUES ($1, $2, $3) RETURNING id", postsListsTable)
	row := tx.QueryRow(createPostQuery, post.Title, post.Description, post.Likes)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PostListPostgres) GetAll() ([]models.Post, error) {
	var posts []models.Post
	query := fmt.Sprintf("SELECT id, title, description, likes FROM %s", postsListsTable)
	err := r.db.Select(&posts, query)

	return posts, err
}