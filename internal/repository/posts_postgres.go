package repository

import (
	"fmt"
	"time"

	"github.com/DavelPurov777/microblog/internal/models"
	"github.com/jmoiron/sqlx"
)

type PostListPostgres struct {
	db *sqlx.DB
}

type postRow struct {
	Id          int       `db:"id"`
	UserId      int       `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	Likes       int       `db:"likes"`
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
	createPostQuery := fmt.Sprintf("INSERT INTO %s (user_id, title, description, created_at, likes) VALUES ($1, $2, $3, $4, $5) RETURNING id", postsListsTable)
	row := tx.QueryRow(createPostQuery, post.UserId, post.Title, post.Description, post.CreatedAt, post.Likes)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PostListPostgres) GetAll() ([]models.Post, error) {
	var rows []postRow
	query := fmt.Sprintf("SELECT id, user_id, title, description, created_at, likes FROM %s", postsListsTable)
	if err := r.db.Select(&rows, query); err != nil {
		return nil, err
	}
	posts := make([]models.Post, len(rows))
	for i := range rows {
		posts[i] = models.Post{
			Id:          rows[i].Id,
			UserId:      rows[i].UserId,
			Title:       rows[i].Title,
			Description: rows[i].Description,
			CreatedAt:   rows[i].CreatedAt,
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
