package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PostLikesRepo struct {
	db *sqlx.DB
}

func NewPostLikesRepo(db *sqlx.DB) *PostLikesRepo {
	return &PostLikesRepo{db: db}
}

func (r *PostLikesRepo) IncrementLikes(userID, postID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.Exec(`
		INSERT INTO likes (user_id, post_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, post_id) DO NOTHING
	`, userID, postID)

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

	_, err = tx.Exec(`
		INSERT INTO engagement_post_likes (post_id, likes)
		VALUES ($1, 1)
		ON CONFLICT (post_id)
		DO UPDATE SET likes = engagement_post_likes.likes + 1
	`, postID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

// TODO: погуглить что делает функция ниже
func isUniqueViolation(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505"
	}

	return false
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
