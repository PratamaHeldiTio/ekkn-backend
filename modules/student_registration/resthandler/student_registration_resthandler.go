package resthandler

import "github.com/gin-gonic/gin"

type StudentRegistrationReshandler interface {
	CreateStudentRegistration(c *gin.Context)
	FindStudentRegistrationByStudentId(c *gin.Context)
}
