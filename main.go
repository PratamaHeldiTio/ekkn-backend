package main

import (
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/modules/student/resthandler"
	"backend-ekkn/modules/student/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=pratama password=mecandoit dbname=ekknutm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	studentRepository := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepository)
	studentReshandler := resthandler.NewStudentResthandler(studentService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/students", studentReshandler.CreateStudent)
	api.POST("/auth/students/login", studentReshandler.LoginStudent)
	api.GET("/students/:nim", studentReshandler.FindStudentByNim)

	router.Run()
}
