package service

import (
	"testing"
)

type MockQueue struct {
	published []LikeEvent
	ch        chan LikeEvent
}

func NewMockQueue() *MockQueue {
	return &MockQueue{ch: make(chan LikeEvent, 10)}
}

func (m *MockQueue) Publish(ev LikeEvent) {
	m.published = append(m.published, ev)
}

func (m *MockQueue) Channel() <-chan LikeEvent {
	return m.ch
}

func TestService_LikePost(t *testing.T) {
	q := &MockQueue{}
	s := &PostListService{queue: q}

	err := s.LikePost(1, 42)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if len(q.published) != 1 || q.published[0].PostID != 1 {
		t.Errorf("publish not called correctly")
	}
}
