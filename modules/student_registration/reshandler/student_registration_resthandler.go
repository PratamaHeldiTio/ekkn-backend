package reshandler

import "github.com/gin-gonic/gin"

type StudentRegistrationReshandler interface {
	CreateStudentRegistration(c *gin.Context)
}
