package controllers

import (
	"context"

	pb "github.com/douglira/go-grpc/proto"
	"github.com/douglira/go-grpc/server/database"
	"github.com/douglira/go-grpc/server/models"
)

type StudentServer struct {
	pb.UnimplementedStudentServer
}

func (s *StudentServer) GetStudent(ctx context.Context, r *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	var student pb.GetStudentResponse
	result := database.DB.Model(&models.Student{}).First(&student, r.StudentId)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &student, nil
}
