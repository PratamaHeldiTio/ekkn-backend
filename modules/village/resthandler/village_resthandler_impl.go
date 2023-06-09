package resthandler

import (
	"backend-ekkn/modules/village/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VillageResthandlerImpl struct {
	service service.VillageService
}

func NewVillageResthandler(service service.VillageService) VillageResthandler {
	return &VillageResthandlerImpl{service}
}

func (handler *VillageResthandlerImpl) CreateVillage(c *gin.Context) {
	var villageRequest shareddomain.RequestVillage
	if err := c.ShouldBindJSON(&villageRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "desa gagal ditambahkan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// call service
	if err := handler.service.CreateVillage(villageRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "desa gagal ditambahkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "desa berhasil ditambahkan")
	c.JSON(http.StatusCreated, response)
}

func (handler *VillageResthandlerImpl) FindByPeriod(c *gin.Context) {
	periodID := c.Param("periodID")
	query := c.Query("search")
	// call service
	villages, err := handler.service.FindVillageByPeriod(periodID, query)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "desa gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// maping data
	villageResponses := []shareddomain.ResponseVillage{}
	for _, village := range villages {
		villageResponses = append(villageResponses, shareddomain.ToVillageResponse(village))
	}

	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "desa berhasil didapatkan", villageResponses)
	c.JSON(http.StatusOK, response)
}

func (handler *VillageResthandlerImpl) UpdateVillage(c *gin.Context) {
	var request shareddomain.UpdateVillageRequest
	ID := c.Param("id")
	request.ID = ID

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "desa gagal diubah", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// call service update
	if err := handler.service.UpdateVillage(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "desa gagal diubah", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	response := helper.APIResponseWithoutData(http.StatusOK, true, "desa berhasil diubah")
	c.JSON(http.StatusOK, response)
}

func (handler *VillageResthandlerImpl) DeleteVillage(c *gin.Context) {
	ID := c.Param("id")

	if err := handler.service.DeleteVillage(ID); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "desa gagal dihapus", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	response := helper.APIResponseWithoutData(http.StatusOK, true, "desa berhasil dihapus")
	c.JSON(http.StatusOK, response)
}

func (handler *VillageResthandlerImpl) FindByID(c *gin.Context) {
	ID := c.Param("id")
	// call service
	village, err := handler.service.FindVillageById(ID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "desa gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	responseData := shareddomain.ToVillageResponse(village)
	response := helper.APIResponseWithData(http.StatusOK, true, "desa berhasil didapatkan", responseData)
	c.JSON(http.StatusOK, response)
}

func (handler *VillageResthandlerImpl) AddDescVillage(c *gin.Context) {
	var request shareddomain.AddDescVillage
	ID := c.Param("id")
	request.ID = ID

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal menambahkan deskripsi", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// call service update
	if err := handler.service.AddDescVillage(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal menambahkan deskripsi", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil menambahkan deskripsi")
	c.JSON(http.StatusOK, response)
}
