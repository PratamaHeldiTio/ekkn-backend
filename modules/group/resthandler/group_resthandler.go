package resthandler

import "github.com/gin-gonic/gin"

type GroupReshandler interface {
	CrateGroup(c *gin.Context)
	JoinGroup(c *gin.Context)
	FindGroupByID(c *gin.Context)
	RegisterGroup(c *gin.Context)
	UpdateGroup(c *gin.Context)
	AddVillage(c *gin.Context)
	UploadReport(c *gin.Context)
	UploadPotentialVillage(c *gin.Context)
	FindByGroupByPeriodLeader(c *gin.Context)
	FindRegisteredGroupByPeriod(c *gin.Context)
	AddLecturer(c *gin.Context)
}
