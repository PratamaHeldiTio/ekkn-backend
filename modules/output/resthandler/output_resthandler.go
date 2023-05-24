package resthandler

import "github.com/gin-gonic/gin"

type OutputResthandler interface {
	CreateOutput(c *gin.Context)
	UpdateOutput(c *gin.Context)
	FindOutputByGroup(c *gin.Context)
}
