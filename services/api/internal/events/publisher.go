package events

type EventPublisher interface {
	PublishPostLiked(PostLikedEvent) error
	PublishPostCreated(PostCreatedEvent) error
}

type UserRegisteredPublisher interface {
	PublishUserRegistered(UserRegisteredEvent) error
}
