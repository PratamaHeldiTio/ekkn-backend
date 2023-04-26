package resthandler

import "github.com/gin-gonic/gin"

type GroupReshandler interface {
	CrateGroup(c *gin.Context)
	JoinGroup(c *gin.Context)
	FindGroupByID(c *gin.Context)
	RegisterGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	AddVillage(c *gin.Context)
	UploadProposal(c *gin.Context)
	UploadReport(c *gin.Context)
}
