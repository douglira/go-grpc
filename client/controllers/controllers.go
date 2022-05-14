package controllers

import (
	"net/http"
	"strconv"

	"github.com/douglira/go-grpc/client/models"
	"github.com/douglira/go-grpc/client/services"
	"github.com/gin-gonic/gin"
)

func AllStudents(c *gin.Context) {
	var students []models.Student

	grpcStudents, err := services.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	if len(grpcStudents.Students) == 0 {
		c.JSON(http.StatusOK, students)
		return
	}

	models.FromProtobufList(&students, grpcStudents)

	c.JSON(200, students)
}

func FindStudentById(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}

	grpcStudent, err := services.GetStudentById(studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if grpcStudent.GetId() == 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	student := &models.Student{}
	student.FromProtobuf(grpcStudent)

	c.JSON(http.StatusOK, student)
}

func RegisterStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	services.RegisterStudent(student)

	c.Writer.WriteHeader(http.StatusCreated)
}
