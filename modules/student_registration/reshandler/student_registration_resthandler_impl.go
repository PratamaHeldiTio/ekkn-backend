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

func (handler *StudentRegistrationResthandlerImpl) FindStudentRegistrationByStudentId(c *gin.Context) {
	// get nim from middleware
	nim := c.MustGet("currentUser").(string)

	// call service
	registeredUsers, err := handler.service.FindStudentRegistrationByStudentID(nim)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusInternalServerError, false, "Riwayat pendaftaran gagal didapatkan", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// map data to response format
	responseRegisteredUsers := []shareddomain.ResponseRegisteredStudent{}
	for _, registeredUser := range registeredUsers {
		responseRegisteredUsers = append(responseRegisteredUsers, shareddomain.ToResponRegiteredStudent(registeredUser))
	}

	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Riwayat pendaftaran berhasil didapatkan", responseRegisteredUsers)
	c.JSON(http.StatusOK, response)
}
