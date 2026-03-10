package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PostLikesRepo struct {
	db *sqlx.DB
}

func NewPostLikesRepo(db *sqlx.DB) *PostLikesRepo {
	return &PostLikesRepo{db: db}
}

func (r *PostLikesRepo) IncrementLikes(postID int, delta int64) error {
	_, err := r.db.Exec(`
		INSERT INTO engagement_post_likes (post_id, likes)
		VALUES ($1, $2)
		ON CONFLICT (post_id)
		DO UPDATE SET likes = engagement_post_likes.likes + EXCLUDED.likes
	`, postID, delta)
	return err
}

func (r *PostLikesRepo) GetLikes(postID int) (int64, error) {
	var likes int64
	err := r.db.QueryRow(`
		SELECT likes
		FROM engagement_post_likes
		WHERE post_id = $1
	`, postID).Scan(&likes)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return likes, err
}
