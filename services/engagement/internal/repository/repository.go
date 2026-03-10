package repository

type PostLikesRepository interface {
	IncrementLikes(postID int, delta int64) error
	GetLikes(postID int) (int64, error)
}
