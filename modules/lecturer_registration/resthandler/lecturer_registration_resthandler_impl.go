package resthandler

import (
	"backend-ekkn/modules/lecturer_registration/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type lecturerRegistrationRestHandlerImpl struct {
	service service.LecturerRegistrationService
}

func NewLecturerRegistrationRestHandler(service service.LecturerRegistrationService) LecturerRegistrationRestHandler {
	return &lecturerRegistrationRestHandlerImpl{service}
}

func (handler *lecturerRegistrationRestHandlerImpl) LecturerRegistration(c *gin.Context) {
	var request shareddomain.LecturerRegistrationRequest
	request.LecturerID = c.MustGet("lecturerID").(string)

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "pendaftaran gagal dilakukan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// call service
	if err := handler.service.LecturerRegistration(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Pendaftaran gagal dilakukan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "Pendaftaran berhasil dilakukan")
	c.JSON(http.StatusCreated, response)
}
