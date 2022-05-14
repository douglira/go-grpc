package services

import (
	"context"
	"time"

	"github.com/douglira/go-grpc/client/integrations"

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
