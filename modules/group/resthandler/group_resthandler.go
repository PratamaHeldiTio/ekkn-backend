package resthandler

import "github.com/gin-gonic/gin"

type GroupReshandler interface {
	CrateGroup(c *gin.Context)
	FindGroupByStudentPeriodID(c *gin.Context)
	JoinGroup(c *gin.Context)
}
