package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/douglira/go-grpc/client/integrations"
	"github.com/douglira/go-grpc/client/models"

	pb "github.com/douglira/go-grpc/proto"
)

func GetAllStudents() (*pb.ListStudents, error) {
	conn := integrations.ServerConnection()
	defer conn.Close()

	pbStudentClient := pb.NewStudentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := pbStudentClient.GetAllStudents(ctx, &pb.Void{})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetStudentById(studentId int) (*pb.Student, error) {
	conn := integrations.ServerConnection()
	defer conn.Close()

	pbStudentClient := pb.NewStudentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := pbStudentClient.GetStudent(ctx, &pb.StudentId{StudentId: int32(studentId)})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func RegisterStudent(student models.Student) {
	p := integrations.NewKafkaProducer()
	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "student-register"
	kafkaValue := new(bytes.Buffer)
	json.NewEncoder(kafkaValue).Encode(student)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          kafkaValue.Bytes(),
	}, nil)

	p.Flush(15 * 1000)
}
