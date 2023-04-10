package main

import (
	"backend-ekkn/jwt_manager"
	"backend-ekkn/middleware"
	"backend-ekkn/migration"
	repository2 "backend-ekkn/modules/period/repository"
	resthandler2 "backend-ekkn/modules/period/resthandler"
	service2 "backend-ekkn/modules/period/service"
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/modules/student/resthandler"
	"backend-ekkn/modules/student/service"
	repository3 "backend-ekkn/modules/student_registration/repository"
	"backend-ekkn/modules/student_registration/reshandler"
	service3 "backend-ekkn/modules/student_registration/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// db connect
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

	// migration
	migration.RunMigration(db)

	// module student
	studentRepository := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepository)
	jwtManager := jwtmanager.NewJwtManager()
	studentReshandler := resthandler.NewStudentResthandler(studentService, jwtManager)

	// module period
	periodRepository := repository2.NewPeriodRepository(db)
	periodService := service2.NewPeriodService(periodRepository)
	periodReshandler := resthandler2.NewPeriodResthandler(periodService)

	// module student registration
	studentRegistrationRepository := repository3.NewStudentRegistrationRepository(db)
	studentRegistrationService := service3.NewStudentRegistrationService(studentRegistrationRepository)
	studentRegistrationResthandler := reshandler.NewStudentRegistrationResthandler(studentRegistrationService)

	// init router gin
	router := gin.Default()

	// config cors allow all origin
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())

	// group router
	api := router.Group("/api/v1")

	// middleware
	authMiddleware := middleware.NewAtuhMiddleware(jwtManager)

	// endpoint student
	api.POST("/students", authMiddleware.AuthMiddleWare(), studentReshandler.CreateStudent)
	api.GET("/students", authMiddleware.AuthMiddleWare(), studentReshandler.FindAllStudent)
	api.POST("/auth/students/login", studentReshandler.LoginStudent)
	api.GET("/students/:nim", studentReshandler.FindStudentByNim)
	api.PUT("/students/:nim", authMiddleware.AuthMiddleWare(), studentReshandler.UpdateStudent)
	api.DELETE("/students/:nim", authMiddleware.AuthMiddleWare(), studentReshandler.DeleteStudent)

	// endpoint period
	api.POST("/period", authMiddleware.AuthMiddleWare(), periodReshandler.CreatePeriod)
	api.PUT("/period", authMiddleware.AuthMiddleWare(), periodReshandler.UpdatePeriod)
	api.GET("/period", authMiddleware.AuthMiddleWare(), periodReshandler.FindAllPeriod)
	api.GET("/period/:id", authMiddleware.AuthMiddleWare(), periodReshandler.FindPeriodById)
	api.DELETE("/period/:id", authMiddleware.AuthMiddleWare(), periodReshandler.DeletePeriodById)

	//endpoint student registration
	api.POST("/student_registration", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.CreateStudentRegistration)

	router.Run()
}
