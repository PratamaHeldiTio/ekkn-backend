package resthandler

import (
	"backend-ekkn/modules/group/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GroupResthandlerImpl struct {
	service service.GroupService
}

func NewGroupReshandler(service service.GroupService) GroupReshandler {
	return &GroupResthandlerImpl{service}
}

func (handler *GroupResthandlerImpl) CrateGroup(c *gin.Context) {
	var requestGroup shareddomain.RequestGroup
	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&requestGroup); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal membuat kelompok", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// asign to struct from param and current user
	requestGroup.PeriodID = c.Param("periodID")
	requestGroup.Leader = c.MustGet("currentUser").(string)

	if err := handler.service.CreateGroup(requestGroup); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal membuat kelompok", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "Berhasil membuat kelompok")
	c.JSON(http.StatusCreated, response)
}

func (handler *GroupResthandlerImpl) FindGroupByStudentPeriodID(c *gin.Context) {
	// get student id from context and period id from param
	PeriodID := c.Param("periodID")
	StudentID := c.MustGet("currentUser").(string)

	group, err := handler.service.FindGroupByStudentPeriodID(StudentID, PeriodID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mendapatkan kelompok", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if group.ID == "" {
		// create response
		response := helper.APIResponseWithError(http.StatusNotFound, false, "Kelompok tidak dapat ditemukan", err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}

	// create response
	responseGroup := shareddomain.ToResponseGroupByStudentPeriodID(group)
	response := helper.APIResponseWithData(http.StatusOK, true, "Berhasil membuat kelompok", responseGroup)
	c.JSON(http.StatusOK, response)
}

func (handler *GroupResthandlerImpl) JoinGroup(c *gin.Context) {
	// get student id from context and period id from param
	PeriodID := c.Param("periodID")
	StudentID := c.MustGet("currentUser").(string)

	var referral shareddomain.RequestJoin
	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&referral); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal bergabung dengan kelompok", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.JoinGroup(StudentID, PeriodID, referral.Referral); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal bergabung dengan kelompok", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil bergabung dengan group")
	c.JSON(http.StatusOK, response)
}
