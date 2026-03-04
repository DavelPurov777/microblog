package models

import "time"

type Post struct {
	Id          int
	UserId      int
	Title       string
	Description string
	CreatedAt   time.Time
	Likes       int
}
