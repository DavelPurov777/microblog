package repository

import (
	"fmt"

	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/jmoiron/sqlx"
)

type PostListPostgres struct {
	db *sqlx.DB
}

type postRow struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Likes       int    `db:"likes"`
}

func NewPostListPostgres(db *sqlx.DB) *PostListPostgres {
	return &PostListPostgres{db: db}
}

func (r *PostListPostgres) Create(post models.Post) (int, error) {
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
	var rows []postRow
	query := fmt.Sprintf("SELECT id, title, description, likes FROM %s", postsListsTable)
	if err := r.db.Select(&rows, query); err != nil {
		return nil, err
	}
	posts := make([]models.Post, len(rows))
	for i := range rows {
		posts[i] = models.Post{
			Id:          rows[i].Id,
			Title:       rows[i].Title,
			Description: rows[i].Description,
			Likes:       rows[i].Likes,
		}
	}

	return posts, nil
}

func (r *PostListPostgres) LikePost(listId int) error {
	query := fmt.Sprintf("UPDATE %s SET likes = likes + 1 WHERE id = $1", postsListsTable)

	_, err := r.db.Exec(query, listId)
	return err
}
