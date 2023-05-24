package resthandler

import (
	"backend-ekkn/modules/output/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OutputRestHandlerImpl struct {
	service service.OutputService
}

func NewOutputRestHandler(service service.OutputService) OutputResthandler {
	return &OutputRestHandlerImpl{service}
}

func (handler *OutputRestHandlerImpl) CreateOutput(c *gin.Context) {
	var request shareddomain.OutputRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal menambahkan luaran", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.CreateOutput(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal menambahkan luaran", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil menambahkan luaran")
	c.JSON(http.StatusOK, response)
}

func (handler *OutputRestHandlerImpl) UpdateOutput(c *gin.Context) {
	var request shareddomain.UpdateOutputRequest
	request.ID = c.Param("id")

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "gagal mengubah luaran", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.UpdateOutput(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mengubah luaran", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil mengubah luaran")
	c.JSON(http.StatusOK, response)
}

func (handler *OutputRestHandlerImpl) FindOutputByGroup(c *gin.Context) {
	ID := c.Param("groupID")

	outputs, err := handler.service.FindOutputByGroup(ID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mengubah luaran", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	outputsData := []shareddomain.OutputResponse{}
	for _, output := range outputs {
		outputsData = append(outputsData, shareddomain.ToOutputResponse(output))
	}

	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "berhasil mendapatkan luaran", outputsData)
	c.JSON(http.StatusOK, response)
}
