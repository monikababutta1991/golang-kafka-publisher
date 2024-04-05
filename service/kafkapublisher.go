package service

import (
	"context"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type kafkaService struct {
	host      string
	topic     string
	partition int
}

type IKafkaService interface {
	ProduceMessage(string)
}

func NewKafkaService(host string, topic string) IKafkaService {
	return &kafkaService{
		host:      host,
		topic:     topic,
		partition: 0,
	}
}

func (k *kafkaService) ProduceMessage(postData string) {

	// to produce messages
	conn, err := kafka.DialLeader(context.Background(), "tcp", k.host, k.topic, k.partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		// return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(postData)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		// return err
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
		// return err
	}
	// return nil
}
