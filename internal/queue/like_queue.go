package queue

import "github.com/DavelPurov777/microblog/internal/service"

type LikeQueue struct {
	ch chan service.LikeEvent
}

func NewLikeQueue(buffer int) *LikeQueue {
	return &LikeQueue{
		ch: make(chan service.LikeEvent, buffer),
	}
}

func (q *LikeQueue) Publish(ev service.LikeEvent) {
	q.ch <- ev
}

func (q *LikeQueue) Channel() <-chan service.LikeEvent {
	return q.ch
}
