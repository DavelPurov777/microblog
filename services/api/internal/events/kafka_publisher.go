package events

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type KafkaLikePublisher struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaLikePublisher(brokers []string, topic, clientID string) *KafkaLikePublisher {
	return &KafkaLikePublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
		topic: topic,
	}
}

func (p *KafkaLikePublisher) PublishLike(ev LikeEvent) error {
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
