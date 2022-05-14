package integrations

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Fatalf("Server not connected: %v", err)
	}
	return p
}
