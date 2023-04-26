package resthandler

import (
	"backend-ekkn/modules/group/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

func (handler *GroupResthandlerImpl) JoinGroup(c *gin.Context) {
	// get student id from context and period id from param
	PeriodID := c.Param("periodID")
	Nim := c.MustGet("currentUser").(string)

	var referral shareddomain.RequestJoin
	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&referral); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal bergabung dengan kelompok", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.JoinGroup(Nim, PeriodID, referral.Referral); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal bergabung dengan kelompok", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil bergabung dengan group")
	c.JSON(http.StatusOK, response)
}

func (handler *GroupResthandlerImpl) FindGroupByID(c *gin.Context) {
	ID := c.Param("id")
	group, err := handler.service.FindGroupID(ID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mendapatkan kelompok", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responseData := shareddomain.ToResponseGroupByID(group)
	response := helper.APIResponseWithData(http.StatusOK, true, "berhasil mendapatkan kelompok", responseData)
	c.JSON(http.StatusOK, response)
}

func (handler *GroupResthandlerImpl) RegisterGroup(c *gin.Context) {
	// get id group param
	ID := c.Param("id")
	Nim := c.MustGet("currentUser").(string)

	if err := handler.service.RegisterGroup(ID, Nim); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mendaftarkan kelompok", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil mendaftarkan kelompok")
	c.JSON(http.StatusOK, response)
}

func (handler *GroupResthandlerImpl) UpdateGroup(c *gin.Context) {
	// get id group param
	ID := c.Param("id")
	Nim := c.MustGet("currentUser").(string)
	request := shareddomain.GroupUpdateRequest{
		ID:  ID,
		Nim: Nim,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal memperbaharui kelompok", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.UpdateGroup(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal memperbaharui kelompok", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil memperbaharui kelompok")
	c.JSON(http.StatusOK, response)
}

func (handler *GroupResthandlerImpl) AddVillage(c *gin.Context) {
	// get id group param
	ID := c.Param("id")
	Nim := c.MustGet("currentUser").(string)
	request := shareddomain.AddVillage{
		ID:  ID,
		Nim: Nim,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal menambahkan desa", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.AddVillage(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal menambahkan desa", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil menambahkan desa")
	c.JSON(http.StatusOK, response)
}

func (handler *GroupResthandlerImpl) UploadProposal(c *gin.Context) {
	file, err := c.FormFile("proposal")
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload proposal", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if file.Size > 10485760 {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload proposal", "file terlalu besar")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if file.Header.Values("Content-Type")[0] != "application/pdf" {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload proposal", "format file salah")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save file to directory
	path := "public/proposal/" + strconv.FormatInt(time.Now().UnixMilli(), 10) + "_" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload proposal", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save path to db
	groupUpdateRequest := shareddomain.GroupUpdateRequest{
		ID:       c.Param("id"),
		Proposal: path,
		Nim:      c.MustGet("currentUser").(string),
	}
	if err := handler.service.UpdateGroup(groupUpdateRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload proposal", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil upload proposal")
	c.JSON(http.StatusOK, response)

}

func (handler *GroupResthandlerImpl) UploadReport(c *gin.Context) {
	file, err := c.FormFile("report")
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload laporan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if file.Size > 10485760 {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload laporan", "file terlalu besar")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if file.Header.Values("Content-Type")[0] != "application/pdf" {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload laporan", "format file salah")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save file to directory
	path := "public/report/" + strconv.FormatInt(time.Now().UnixMilli(), 10) + "_" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload laporan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// save path to db
	groupUpdateRequest := shareddomain.GroupUpdateRequest{
		ID:     c.Param("id"),
		Report: path,
		Nim:    c.MustGet("currentUser").(string),
	}
	if err := handler.service.UpdateGroup(groupUpdateRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal upload proposal", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil upload proposal")
	c.JSON(http.StatusOK, response)

}
