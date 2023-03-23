package main

import (
	"backend-ekkn/jwt_manager"
	"backend-ekkn/middleware"
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
	jwtManager := jwtmanager.NewJwtManager()
	studentReshandler := resthandler.NewStudentResthandler(studentService, jwtManager)

	router := gin.Default()
	api := router.Group("/api/v1")

	authMiddleware := middleware.NewAtuhMiddleware(jwtManager)

	api.POST("/students", authMiddleware.AuthMiddleWare(), studentReshandler.CreateStudent)
	api.GET("/students", authMiddleware.AuthMiddleWare(), studentReshandler.FindAllStudent)
	api.POST("/auth/students/login", studentReshandler.LoginStudent)
	api.GET("/students/:nim", authMiddleware.AuthMiddleWare(), studentReshandler.FindStudentByNim)
	api.PUT("/students/:nim", authMiddleware.AuthMiddleWare(), studentReshandler.UpdateStudent)

	router.Run()
}
