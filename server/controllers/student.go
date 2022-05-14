package controllers

import (
	"context"

	pb "github.com/douglira/go-grpc/proto"
	"github.com/douglira/go-grpc/server/database"
	"github.com/douglira/go-grpc/server/models"
)

type StudentServer struct {
	pb.UnimplementedStudentServiceServer
}

func (s *StudentServer) GetStudent(ctx context.Context, r *pb.StudentId) (*pb.Student, error) {
	var student pb.Student
	result := database.DB.Model(&models.Student{}).First(&student, r.StudentId)
	if result.RowsAffected == 0 {
		return &student, nil
	}
	return &student, nil
}

func (s *StudentServer) GetAllStudents(ctx context.Context, r *pb.Void) (*pb.ListStudents, error) {
	var listStudent pb.ListStudents = pb.ListStudents{Students: []*pb.Student{}}
	result := database.DB.Model(&models.Student{}).Find(&listStudent.Students)
	if result.RowsAffected == 0 {
		return &listStudent, nil
	}
	return &listStudent, nil
}
