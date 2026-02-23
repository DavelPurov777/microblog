package queue

type LikeQueue struct {
	ch chan int
}

func NewLikeQueue(buffer int) *LikeQueue {
	return &LikeQueue{
		ch: make(chan int, buffer),
	}
}

func (q *LikeQueue) Publish(id int) {
	q.ch <- id
}

func (q *LikeQueue) Channel() <-chan int {
	return q.ch
}
