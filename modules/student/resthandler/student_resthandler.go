package resthandler

import "github.com/gin-gonic/gin"

type StudentResthandler interface {
	CreateStudent(c *gin.Context)
}
