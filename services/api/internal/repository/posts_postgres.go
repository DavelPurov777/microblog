package repository

import (
	"time"

	"github.com/DavelPurov777/microblog/services/api/internal/models"
	sq "github.com/Masterminds/squirrel"
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
	query, args, err := psql.
		Insert(postsListsTable).
		Columns("user_id", "title", "description", "created_at", "likes").
		Values(post.UserId, post.Title, post.Description, post.CreatedAt, post.Likes).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	row := tx.QueryRow(query, args...)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PostListPostgres) GetAll() ([]models.Post, error) {
	var rows []postRow
	query, args, err := psql.
		Select("id", "user_id", "title", "description", "created_at", "likes").
		From(postsListsTable).
		ToSql()
	if err != nil {
		return nil, err
	}

	if err := r.db.Select(&rows, query, args...); err != nil {
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

func (r *PostListPostgres) LikePost(postId, userId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	insertQ, insertArgs, err := psql.
		Insert("likes").
		Columns("user_id", "post_id").
		Values(userId, postId).
		Suffix("ON CONFLICT (user_id, post_id) DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}

	res, err := tx.Exec(insertQ, insertArgs...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return tx.Commit()
	}

	updateQ, updateArgs, err := psql.
		Update(postsListsTable).
		Set("likes", sq.Expr("likes + 1")).
		Where(sq.Eq{"id": postId}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(updateQ, updateArgs...)
	if err != nil {
		return err
	}

	return tx.Commit()
}
