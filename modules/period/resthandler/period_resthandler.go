package resthandler

import "github.com/gin-gonic/gin"

type PeriodResthandler interface {
	CreatePeriod(c *gin.Context)
}
