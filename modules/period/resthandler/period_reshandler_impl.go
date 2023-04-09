package resthandler

import (
	"backend-ekkn/modules/period/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type PeriodResthandlerImpl struct {
	service service.PeriodService
}

// init resthandler
func NewPeriodResthandler(service service.PeriodService) PeriodResthandler {
	return &PeriodResthandlerImpl{service}
}

// handler create period
func (handler *PeriodResthandlerImpl) CreatePeriod(c *gin.Context) {
	// validation request
	var periodRequest shareddomain.RequestPeriod
	if err := c.ShouldBindJSON(&periodRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Periode gagal ditambahkan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send request to service
	if err := handler.service.CreatePeriod(periodRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Periode gagal ditambahkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "Periode berhasil ditambahkan")
	c.JSON(http.StatusCreated, response)

}

func (handler *PeriodResthandlerImpl) FindAllPeriod(c *gin.Context) {
	periods, err := handler.service.FindAllPeriod()
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusInternalServerError, false, "Periode gagal didapatkan", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get role
	curentRole := c.MustGet("currentRole")

	if curentRole == "student" || curentRole == "lecture" {
		// init type of response
		responsePeriods := []shareddomain.ResponsePeriodBasic{}

		// maping data
		for _, period := range periods {
			responsePeriods = append(responsePeriods, shareddomain.ToResponsePeriodBasic(period))
		}

		// create response
		response := helper.APIResponseWithData(http.StatusOK, true, "Periode berhasil didapatkan", responsePeriods)
		c.JSON(http.StatusOK, response)
		return
	}

	// init type of response
	responsePeriods := []shareddomain.ResponsePeriod{}

	// maping data
	for _, period := range periods {
		responsePeriods = append(responsePeriods, shareddomain.ToResponsePeriod(period))
	}
	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Periode berhasil didapatkan", responsePeriods)
	c.JSON(http.StatusOK, response)
}

func (handler *PeriodResthandlerImpl) FindPeriodById(c *gin.Context) {
	// get id with params url
	idString := c.Param("id")

	//parse uuid string to uuid
	id, err := uuid.Parse(idString)
	if err != nil {
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Periode gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	period, err := handler.service.FindPeriodById(id)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Periode gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// maping data
	responsePeriod := shareddomain.ToResponsePeriod(period)

	// create response success
	response := helper.APIResponseWithData(http.StatusOK, true, "Periode berhasil didapatkan", responsePeriod)
	c.JSON(http.StatusOK, response)
}

// handler update

func (handler *PeriodResthandlerImpl) UpdatePeriod(c *gin.Context) {
	// validation request
	var periodRequest shareddomain.RequestPeriod
	if err := c.ShouldBindJSON(&periodRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Periode gagal diubah", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// call service update
	if err := handler.service.UpdatePeriod(periodRequest); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Periode gagal diubah", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	response := helper.APIResponseWithoutData(http.StatusOK, true, "Periode berhasil diubah")
	c.JSON(http.StatusOK, response)

}

// handler delete period

func (handler *PeriodResthandlerImpl) DeletePeriodById(c *gin.Context) {
	// get id with params url
	idString := c.Param("id")

	//parse uuid string to uuid
	id, err := uuid.Parse(idString)
	if err != nil {
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Periode gagal dihapuskan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := handler.service.DeletePeriodById(id); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Periode gagal dihapuskan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseWithoutData(http.StatusOK, true, "Periode berhasil dihapuskan")
	c.JSON(http.StatusOK, response)
}
