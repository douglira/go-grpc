package models

import pb "github.com/douglira/go-grpc/proto"

type Student struct {
	Id                        int    `json:"id,omitempty"`
	Name                      string `json:"name,omitempty"`
	IdentityNumber            string `json:"identityNumber,omitempty"`
	GeneralRegistrationNumber string `json:"generalRegistrationNumber,omitempty"`
}

func FromProtobufList(students *[]Student, list *pb.ListStudents) {
	for _, spb := range list.Students {
		var student Student
		student.FromProtobuf(spb)
		*students = append(*students, student)
	}
}

func (s *Student) FromProtobuf(spb *pb.Student) {
	s.Id = int(spb.GetId())
	s.Name = spb.GetName()
	s.IdentityNumber = spb.GetIdentityNumber()
	s.GeneralRegistrationNumber = spb.GetGeneralRegistrationNumber()
}
