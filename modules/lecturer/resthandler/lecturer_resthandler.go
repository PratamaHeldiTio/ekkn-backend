package resthandler

import "github.com/gin-gonic/gin"

type LecturerRestHandler interface {
	CreateLecturer(c *gin.Context)
	UpdateLecturer(c *gin.Context)
	UpdateLecturerByJwt(c *gin.Context)
	FindByIdParam(c *gin.Context)
	FindByIdJwt(c *gin.Context)
	FindAllLecturer(c *gin.Context)
	LoginLecturer(c *gin.Context)
	DeleteLecturer(c *gin.Context)
	ResetPassword(c *gin.Context)
	ChangePassword(c *gin.Context)
}
