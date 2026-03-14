package events

type EventPublisher interface {
	PublishPostLiked(PostLikedEvent) error
	PublishPostCreated(PostCreatedEvent) error
	PublishUserRegistered(UserRegisteredEvent) error
}
