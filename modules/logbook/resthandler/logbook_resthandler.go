package resthandler

import "github.com/gin-gonic/gin"

type LogbookRestHandler interface {
	CreateLogbook(c *gin.Context)
	FindLogbookByStudentPeriod(c *gin.Context)
	FindLogbookByStudentPeriodParam(c *gin.Context)
}
