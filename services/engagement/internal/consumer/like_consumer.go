package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/DavelPurov777/microblog/services/engagement/internal/events"
	"github.com/DavelPurov777/microblog/services/engagement/internal/repository"
	"github.com/segmentio/kafka-go"
)

type LikeConsumer struct {
	reader *kafka.Reader
	repo   repository.PostLikesRepository
}

func NewLikeConsumer(brokers []string, topic, groupID string, repo repository.PostLikesRepository) *LikeConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})

	return &LikeConsumer{
		reader: r,
		repo:   repo,
	}
}

func (c *LikeConsumer) Run(ctx context.Context) error {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		var ev events.LikeEvent
		if err := json.Unmarshal(m.Value, &ev); err != nil {
			log.Printf("failed to unmarshal like event: %v", err)
			continue
		}

		if err := c.repo.IncrementLikes(ev.UserID, ev.PostID); err != nil {
			log.Printf("failed to increment likes %v", err)
			continue
		}
	}
}
