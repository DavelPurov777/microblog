package events

import "time"

type LikeEvent struct {
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
