package repository

type PostLikesRepository interface {
	IncrementLikes(userID, postID int) error
	GetLikes(postID int) (int64, error)
}
