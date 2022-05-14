package routes

import (
	"github.com/douglira/go-grpc/client/controllers"
	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine) {
	apiRouter := r.Group("/api")
	apiRouter.GET("/students", controllers.AllStudents)
	apiRouter.POST("/students", controllers.RegisterStudent)
	apiRouter.GET("/students/:id", controllers.FindStudentById)
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	router(r)
	return r
}
