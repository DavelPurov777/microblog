package service

import (
	"testing"
)

type MockQueue struct {
	published []int
	ch        chan int
}

func (m *MockQueue) Publish(id int) {
	m.published = append(m.published, id)
}

func (m *MockQueue) Channel() <-chan int {
	return m.ch
}

func TestService_LikePost(t *testing.T) {
	q := &MockQueue{}
	s := &PostListService{queue: q}

	err := s.LikePost(42)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	if len(q.published) != 1 || q.published[0] != 42 {
		t.Errorf("publish not called correctly")
	}
}
