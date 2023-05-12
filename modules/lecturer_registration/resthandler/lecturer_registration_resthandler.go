package resthandler

import "github.com/gin-gonic/gin"

type LecturerRegistrationRestHandler interface {
	LecturerRegistration(c *gin.Context)
}
