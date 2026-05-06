package events

type PostsEventPublisher interface {
	PublishPostLiked(PostLikedEvent) error
	PublishPostCreated(PostCreatedEvent) error
}

type UserEventPublisher interface {
	PublishUserRegistered(UserRegisteredEvent) error
}
