package events

type LikeEventPublisher interface {
	PublishLike(LikeEvent) error
}
