package handlers

import (
	"bytes"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/douglira/go-grpc/server/database"
	"github.com/douglira/go-grpc/server/integrations"
	"github.com/douglira/go-grpc/server/models"
)

type studentRegisterDlq struct {
	MessageValue []byte
	Error        error
}

func saveStudent(student *models.Student) error {
	err := database.DB.Model(models.Student{}).Save(student)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func getDlqStudentRegisterPayload(err error, msg []byte) studentRegisterDlq {
	return studentRegisterDlq{
		MessageValue: msg,
		Error:        err,
	}
}

func registerStudentTopicSubscription() {
	dlqTopic := "dql.student-register"
	dlqChannel := make(chan studentRegisterDlq)

	go func() {
		c := integrations.NewKafkaConsumer()
		topic := "student-register"
		c.Subscribe(topic, nil)
		defer c.Close()
		for {
			msg, err := c.ReadMessage(-1)
			if err != nil {
				dlqChannel <- getDlqStudentRegisterPayload(err, msg.Value)
			} else {
				var student models.Student
				err := json.Unmarshal(msg.Value, &student)
				if err != nil {
					dlqChannel <- getDlqStudentRegisterPayload(err, msg.Value)
					return
				}
				err = saveStudent(&student)
				if err != nil {
					dlqChannel <- getDlqStudentRegisterPayload(err, msg.Value)
				}
			}
		}
	}()

	go func() {
		p := integrations.NewKafkaProducer()
		defer p.Close()

		for dlqValue := range dlqChannel {
			dlqMsg := new(bytes.Buffer)
			json.NewEncoder(dlqMsg).Encode(dlqValue)
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &dlqTopic, Partition: kafka.PartitionAny},
				Value:          dlqMsg.Bytes(),
			}, nil)
		}

	}()

}

func ExecuteKafkaHandlers() {
	registerStudentTopicSubscription()
}
