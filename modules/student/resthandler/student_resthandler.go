package resthandler

import "github.com/gin-gonic/gin"

type StudentResthandler interface {
	CreateStudent(c *gin.Context)
	FindStudentByNim(c *gin.Context)
	LoginStudent(c *gin.Context)
	FindAllStudent(c *gin.Context)
	UpdateStudent(c *gin.Context)
	DeleteStudent(c *gin.Context)
	ChangePassword(c *gin.Context)
}
