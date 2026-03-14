package events

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type KafkaEventPublisher struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaPublisher(brokers []string, topic, clientID string) *KafkaEventPublisher {
	return &KafkaEventPublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
		topic: topic,
	}
}

func (p *KafkaEventPublisher) PublishPostLiked(ev PostLikedEvent) error {
	data, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(strconv.Itoa(ev.PostID)),
		Value: data,
	}

	return p.writer.WriteMessages(context.Background(), msg)
}

func (p *KafkaEventPublisher) PublishPostCreated(ev PostCreatedEvent) error {
	data, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(strconv.Itoa(ev.Id)),
		Value: data,
	}

	return p.writer.WriteMessages(context.Background(), msg)
}

func (p *KafkaEventPublisher) PublishUserRegistered(ev UserRegisteredEvent) error {
	data, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(strconv.Itoa(ev.Id)),
		Value: data,
	}

	return p.writer.WriteMessages(context.Background(), msg)
}
