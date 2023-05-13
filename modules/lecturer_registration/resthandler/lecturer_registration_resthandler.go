package resthandler

import (
	"github.com/gin-gonic/gin"
)

type LecturerRegistrationRestHandler interface {
	LecturerRegistration(c *gin.Context)
	FindLecturerRegistrationHistory(c *gin.Context)
	ValidationLecturerRegistration(c *gin.Context)
	FindLecturerRegistrationByPeriod(c *gin.Context)
}
