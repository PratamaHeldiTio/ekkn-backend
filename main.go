package main

import (
	"backend-ekkn/jwt_manager"
	"backend-ekkn/middleware"
	"backend-ekkn/migration"
	repository4 "backend-ekkn/modules/group/repository"
	resthandler4 "backend-ekkn/modules/group/resthandler"
	service4 "backend-ekkn/modules/group/service"
	repository2 "backend-ekkn/modules/period/repository"
	resthandler2 "backend-ekkn/modules/period/resthandler"
	service2 "backend-ekkn/modules/period/service"
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/modules/student/resthandler"
	"backend-ekkn/modules/student/service"
	repository3 "backend-ekkn/modules/student_registration/repository"
	resthandler3 "backend-ekkn/modules/student_registration/resthandler"
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
	studentResthandler := resthandler.NewStudentResthandler(studentService, jwtManager)

	// module period
	periodRepository := repository2.NewPeriodRepository(db)
	periodService := service2.NewPeriodService(periodRepository)
	periodResthandler := resthandler2.NewPeriodResthandler(periodService)

	// module student registration
	studentRegistrationRepository := repository3.NewStudentRegistrationRepository(db)
	studentRegistrationService := service3.NewStudentRegistrationService(studentRegistrationRepository)
	studentRegistrationResthandler := resthandler3.NewStudentRegistrationResthandler(studentRegistrationService)

	//module group
	groupRepository := repository4.NewGroupRepository(db)
	groupService := service4.NewGroupServiceImpl(groupRepository, studentRegistrationService)
	groupResthandler := resthandler4.NewGroupReshandler(groupService)

	// init router gin
	router := gin.Default()

	// ini cors middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// add header to allow header config
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	// group router
	api := router.Group("/api/v1")

	// middleware
	authMiddleware := middleware.NewAtuhMiddleware(jwtManager)

	// endpoint student
	api.POST("/student", authMiddleware.AuthMiddleWare(), studentResthandler.CreateStudent)
	api.GET("/student", authMiddleware.AuthMiddleWare(), studentResthandler.FindAllStudent)
	api.POST("/auth/student/login", studentResthandler.LoginStudent)
	api.GET("/student/:nim", authMiddleware.AuthMiddleWare(), studentResthandler.FindStudentByNim)
	api.PUT("/student/:nim", authMiddleware.AuthMiddleWare(), studentResthandler.UpdateStudent)
	api.DELETE("/student/:nim", authMiddleware.AuthMiddleWare(), studentResthandler.DeleteStudent)

	// endpoint period
	api.POST("/period", authMiddleware.AuthMiddleWare(), periodResthandler.CreatePeriod)
	api.PUT("/period", authMiddleware.AuthMiddleWare(), periodResthandler.UpdatePeriod)
	api.GET("/period", authMiddleware.AuthMiddleWare(), periodResthandler.FindAllPeriod)
	api.GET("/period/student", authMiddleware.AuthMiddleWare(), periodResthandler.FindAllPeriodByStudent)
	api.GET("/period/:id", authMiddleware.AuthMiddleWare(), periodResthandler.FindPeriodById)
	api.DELETE("/period/:id", authMiddleware.AuthMiddleWare(), periodResthandler.DeletePeriodById)

	//endpoint student registration
	api.POST("/student/registration", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.CreateStudentRegistration)
	api.GET("/student/registration", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationByStudentId)
	api.GET("/student/registered", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationRegistered)
	api.GET("/student/registration/:periodID", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationByNimPeriodID)

	//endpoint group
	api.POST("/group/:periodID", authMiddleware.AuthMiddleWare(), groupResthandler.CrateGroup)
	api.GET("/group/:id", authMiddleware.AuthMiddleWare(), groupResthandler.FindGroupByID)
	api.POST("/group/join/:periodID", authMiddleware.AuthMiddleWare(), groupResthandler.JoinGroup)

	router.Run()
}
