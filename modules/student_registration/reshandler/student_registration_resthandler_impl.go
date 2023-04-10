package reshandler

import (
	"backend-ekkn/modules/student_registration/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentRegistrationResthandlerImpl struct {
	service service.StudentRegistrationService
}

// init repo
func NewStudentRegistrationResthandler(service service.StudentRegistrationService) StudentRegistrationReshandler {
	return &StudentRegistrationResthandlerImpl{service}
}

func (handler *StudentRegistrationResthandlerImpl) CreateStudentRegistration(c *gin.Context) {
	// get id
	nim := c.MustGet("currentUser").(string)
	// validation request
	requestStudentRegistration := shareddomain.RequestStudentRegistration{
		Nim: nim,
	}
	if err := c.ShouldBindJSON(&requestStudentRegistration); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Pendaftaran gagal dilakukan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// call service
	if err := handler.service.CreateStudentRegistration(requestStudentRegistration); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Pendaftaran gagal dilakukan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "Pendaftaran berhasil dilakukan")
	c.JSON(http.StatusCreated, response)
}
