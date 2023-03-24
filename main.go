package main

import (
	"backend-ekkn/jwt_manager"
	"backend-ekkn/middleware"
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/modules/student/resthandler"
	"backend-ekkn/modules/student/service"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("USER_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
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
	api.DELETE("/students/:nim", authMiddleware.AuthMiddleWare(), studentReshandler.DeleteStudent)

	router.Run()
}
