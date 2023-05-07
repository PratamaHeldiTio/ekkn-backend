package resthandler

import "github.com/gin-gonic/gin"

type VillageResthandler interface {
	CreateVillage(c *gin.Context)
	FindByPeriod(c *gin.Context)
	FindByID(c *gin.Context)
	UpdateVillage(c *gin.Context)
	DeleteVillage(c *gin.Context)
	AddDescVillage(c *gin.Context)
}
