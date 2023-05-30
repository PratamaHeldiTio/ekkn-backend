package resthandler

import "github.com/gin-gonic/gin"

type AdminRestHandler interface {
	CreateAdmin(c *gin.Context)
	LoginAdmin(c *gin.Context)
	DeleteAdmin(c *gin.Context)
}
