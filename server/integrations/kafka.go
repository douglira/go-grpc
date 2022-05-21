package integrations

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaConsumer() *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "student-worker",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Server not connected: %v", err)
	}
	return c
}

func NewKafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Fatalf("Server not connected: %v", err)
	}
	return p
}
