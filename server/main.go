package main

import (
	"log"
	"net"

	pb "github.com/douglira/go-grpc/proto"
	"github.com/douglira/go-grpc/server/controllers"
	"github.com/douglira/go-grpc/server/database"
	"google.golang.org/grpc"
)

func main() {
	database.InitiateConnection()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterStudentServer(s, &controllers.StudentServer{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
