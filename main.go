package main

import (
	"backend-ekkn/jwt_manager"
	"backend-ekkn/middleware"
	"backend-ekkn/migration"
	repository7 "backend-ekkn/modules/admin/repository"
	resthandler7 "backend-ekkn/modules/admin/resthandler"
	service7 "backend-ekkn/modules/admin/service"
	repository4 "backend-ekkn/modules/group/repository"
	resthandler4 "backend-ekkn/modules/group/resthandler"
	service4 "backend-ekkn/modules/group/service"
	repository8 "backend-ekkn/modules/lecturer/repository"
	resthandler8 "backend-ekkn/modules/lecturer/resthandler"
	service8 "backend-ekkn/modules/lecturer/service"
	repository6 "backend-ekkn/modules/logbook/repository"
	resthandler6 "backend-ekkn/modules/logbook/resthandler"
	service6 "backend-ekkn/modules/logbook/service"
	repository2 "backend-ekkn/modules/period/repository"
	resthandler2 "backend-ekkn/modules/period/resthandler"
	service2 "backend-ekkn/modules/period/service"
	"backend-ekkn/modules/student/repository"
	"backend-ekkn/modules/student/resthandler"
	"backend-ekkn/modules/student/service"
	repository3 "backend-ekkn/modules/student_registration/repository"
	resthandler3 "backend-ekkn/modules/student_registration/resthandler"
	service3 "backend-ekkn/modules/student_registration/service"
	repository5 "backend-ekkn/modules/village/repository"
	resthandler5 "backend-ekkn/modules/village/resthandler"
	service5 "backend-ekkn/modules/village/service"
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

	// module villge
	villageRepository := repository5.NeWVillageRepository(db)
	villageService := service5.NewVillageService(villageRepository)
	villageResthandler := resthandler5.NewVillageResthandler(villageService)

	//module group
	groupRepository := repository4.NewGroupRepository(db)
	groupService := service4.NewGroupServiceImpl(groupRepository, studentRegistrationService, villageService)
	groupResthandler := resthandler4.NewGroupReshandler(groupService)

	//module logbook
	logbookRepository := repository6.NewLogbookRepository(db)
	logbookService := service6.NewLogbookService(logbookRepository, periodService, groupService)
	logbookRestHandler := resthandler6.NewLogbookResthandler(logbookService)

	// module admin
	adminRespository := repository7.NewAdminRepository(db)
	adminService := service7.NewAdminRepository(adminRespository)
	adminRestHandler := resthandler7.NewAdminRestHandler(adminService, jwtManager)

	// module lecturer
	lecturerRepository := repository8.NewLectureRepositoryImpl(db)
	lecturerService := service8.NewLecturerService(lecturerRepository)
	lecturerRestHandler := resthandler8.NewLecturerRestHandler(lecturerService, jwtManager)

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
	router.Static("/static", "./public")

	// middleware
	authMiddleware := middleware.NewAtuhMiddleware(jwtManager)

	// endpoint student
	api.POST("/student", authMiddleware.AuthMiddleWareAdmin(), studentResthandler.CreateStudent)
	api.GET("/students", authMiddleware.AuthMiddleWareAdmin(), studentResthandler.FindAllStudent)
	api.POST("/auth/student/login", studentResthandler.LoginStudent)
	api.GET("/student", authMiddleware.AuthMiddleWare(), studentResthandler.FindStudentByNim)
	api.GET("/student/:studentID", authMiddleware.AuthMiddleWare(), studentResthandler.FindStudentByNimParam)
	api.PUT("/student", authMiddleware.AuthMiddleWare(), studentResthandler.UpdateStudent)
	api.PUT("/student/:studentID", authMiddleware.AuthMiddleWareAdmin(), studentResthandler.UpdateStudentIDParam)
	api.DELETE("/student/:nim", authMiddleware.AuthMiddleWareAdmin(), studentResthandler.DeleteStudent)
	api.PUT("/student/change_password", authMiddleware.AuthMiddleWare(), studentResthandler.ChangePassword)
	api.PUT("/student/reset_password/:studentID", authMiddleware.AuthMiddleWareAdmin(), studentResthandler.ResetPassword)

	// endpoint period
	api.POST("/period", authMiddleware.AuthMiddleWareAdmin(), periodResthandler.CreatePeriod)
	api.PUT("/period/:id", authMiddleware.AuthMiddleWareAdmin(), periodResthandler.UpdatePeriod)
	api.GET("/period", authMiddleware.AuthMiddleWare(), periodResthandler.FindAllPeriod)
	api.GET("/period/student", authMiddleware.AuthMiddleWare(), periodResthandler.FindAllPeriodByStudent)
	api.GET("/period/:id", authMiddleware.AuthMiddleWare(), periodResthandler.FindPeriodById)
	api.DELETE("/period/:id", authMiddleware.AuthMiddleWareAdmin(), periodResthandler.DeletePeriodById)

	//endpoint student registration
	api.POST("/student/registration", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.CreateStudentRegistration)
	api.GET("/student/registration", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationByStudentId)
	api.GET("/student/registered", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationRegistered)
	api.GET("/student/registration/:periodID", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationByNimPeriodID)
	api.GET("/student/registration/:periodID/:studentID", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationByNimPeriodIDParams)
	api.GET("/student/registrations/:periodID", authMiddleware.AuthMiddleWare(), studentRegistrationResthandler.FindStudentRegistrationByPeriod)
	api.PUT("/student/registration/:ID", authMiddleware.AuthMiddleWareAdmin(), studentRegistrationResthandler.UpdateStudentRegistration)

	//endpoint group
	api.POST("/group/:periodID", authMiddleware.AuthMiddleWare(), groupResthandler.CrateGroup)
	api.GET("/group/:id", authMiddleware.AuthMiddleWare(), groupResthandler.FindGroupByID)
	api.POST("/group/join/:periodID", authMiddleware.AuthMiddleWare(), groupResthandler.JoinGroup)
	api.PUT("/group/register/:id", authMiddleware.AuthMiddleWare(), groupResthandler.RegisterGroup)
	api.POST("/group/village/:id", authMiddleware.AuthMiddleWare(), groupResthandler.AddVillage)
	api.POST("/group/report/:id", authMiddleware.AuthMiddleWare(), groupResthandler.UploadReport)
	api.POST("/group/potential/:id", authMiddleware.AuthMiddleWare(), groupResthandler.UploadPotentialVillage)
	api.GET("/group/leader/:periodID", authMiddleware.AuthMiddleWare(), groupResthandler.FindByGroupByPeriodLeader)

	// endpoint village
	api.POST("/village", authMiddleware.AuthMiddleWareAdmin(), villageResthandler.CreateVillage)
	api.GET("/village/period/:periodID", authMiddleware.AuthMiddleWare(), villageResthandler.FindByPeriod)
	api.GET("/village/:id", authMiddleware.AuthMiddleWare(), villageResthandler.FindByID)
	api.PUT("/village/:id", authMiddleware.AuthMiddleWareAdmin(), villageResthandler.UpdateVillage)
	api.PUT("/village/add_desc/:id", authMiddleware.AuthMiddleWare(), villageResthandler.AddDescVillage)
	api.DELETE("/village/:id", authMiddleware.AuthMiddleWareAdmin(), villageResthandler.DeleteVillage)

	// logbook
	api.POST("/logbook", authMiddleware.AuthMiddleWare(), logbookRestHandler.CreateLogbook)
	api.GET("/logbooks/:periodID", authMiddleware.AuthMiddleWare(), logbookRestHandler.FindLogbookByStudentPeriod)
	api.GET("/logbook/:periodID/:studentID", authMiddleware.AuthMiddleWare(), logbookRestHandler.FindLogbookByStudentPeriodParam)

	// admin
	api.POST("/admin", authMiddleware.AuthMiddleWareAdmin(), adminRestHandler.CreateAdmin)
	api.POST("/auth/admin/login", adminRestHandler.LoginAdmin)

	// lecturer
	api.POST("/lecturer", authMiddleware.AuthMiddleWareAdmin(), lecturerRestHandler.CreateLecturer)
	api.PUT("/lecturer/:id", authMiddleware.AuthMiddleWareAdmin(), lecturerRestHandler.UpdateLecturer)
	api.DELETE("/lecturer/:id", authMiddleware.AuthMiddleWareAdmin(), lecturerRestHandler.DeleteLecturer)
	api.PUT("/lecturer/reset_password/:id", authMiddleware.AuthMiddleWareAdmin(), lecturerRestHandler.ResetPassword)
	api.GET("/lecturer/:id", authMiddleware.AuthMiddleWareAdmin(), lecturerRestHandler.FindByIdParam)
	api.GET("/lecturer", authMiddleware.AuthMiddleWareLecturer(), lecturerRestHandler.FindByIdJwt)
	api.GET("/lecturers", authMiddleware.AuthMiddleWareAdmin(), lecturerRestHandler.FindAllLecturer)
	api.POST("/auth/lecturer/login", lecturerRestHandler.LoginLecturer)

	router.Run()
}
