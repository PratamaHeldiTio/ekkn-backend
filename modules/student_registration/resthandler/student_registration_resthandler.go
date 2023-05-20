package resthandler

import "github.com/gin-gonic/gin"

type StudentRegistrationReshandler interface {
	CreateStudentRegistration(c *gin.Context)
	FindStudentRegistrationByStudentId(c *gin.Context)
	FindStudentRegistrationRegistered(c *gin.Context)
	FindStudentRegistrationByNimPeriodID(c *gin.Context)
	FindStudentRegistrationByNimPeriodIDParams(c *gin.Context)
	FindStudentRegistrationByPeriod(c *gin.Context)
	UpdateStudentRegistration(c *gin.Context)
	AddProkerStudent(c *gin.Context)
	FindStudentRegistrationByID(c *gin.Context)
	FindStudentRegistrationByGroup(c *gin.Context)
}
