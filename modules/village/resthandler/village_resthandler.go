package resthandler

import "github.com/gin-gonic/gin"

type VillageResthandler interface {
	CreateVillage(c *gin.Context)
	FindAllVillage(c *gin.Context)
}
