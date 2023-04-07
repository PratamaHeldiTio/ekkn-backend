package resthandler

import (
	"backend-ekkn/modules/period/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
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

func (handler PeriodResthandlerImpl) CreatePeriod(c *gin.Context) {
	// validation request
	var periodRequest shareddomain.RequestPeriod
	if err := c.ShouldBindJSON(&periodRequest); err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusUnprocessableEntity, false, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send request to service
	if err := handler.service.CreatePeriod(periodRequest); err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusBadRequest, false, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	response := helper.APIResponseWithoutData(http.StatusCreated, true)
	c.JSON(http.StatusCreated, response)

}
