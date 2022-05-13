package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/douglira/go-grpc/client/service"
	pb "github.com/douglira/go-grpc/proto"
)

func getStudentById() {
	conn := service.ServerConnection()
	defer conn.Close()

	pbStudentClient := pb.NewStudentClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := pbStudentClient.GetStudent(ctx, &pb.GetStudentRequest{StudentId: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)
}

func main() {
	getStudentById()
}
