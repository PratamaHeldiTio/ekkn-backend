package resthandler

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
	responseRegisteredUsers := []shareddomain.ResponseStudentRegistrationByNim{}
	for _, registeredUser := range registeredUsers {
		responseRegisteredUsers = append(responseRegisteredUsers, shareddomain.ToResponRegiteredStudent(registeredUser))
	}

	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Riwayat pendaftaran berhasil didapatkan", responseRegisteredUsers)
	c.JSON(http.StatusOK, response)
}

func (handler *StudentRegistrationResthandlerImpl) FindStudentRegistrationRegistered(c *gin.Context) {
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
	responseRegisteredUsers := []shareddomain.ResponseRegisteredStudents{}
	for _, registeredUser := range registeredUsers {
		if registeredUser.Status == "true" {
			responseRegisteredUsers = append(responseRegisteredUsers, shareddomain.ToResponseRegisteredStudents(registeredUser))
		}
	}

	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Riwayat pendaftaran berhasil didapatkan", responseRegisteredUsers)
	c.JSON(http.StatusOK, response)
}

func (handler *StudentRegistrationResthandlerImpl) FindStudentRegistrationByNimPeriodID(c *gin.Context) {
	// get nim from middleware
	nim := c.MustGet("currentUser").(string)
	periodID := c.Param("periodID")

	// call service
	registrationStudent, err := handler.service.FindStudentRegistrationByNimPeriodID(nim, periodID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Pendaftaran mahasiswa gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if registrationStudent.ID == "" {
		// create response
		response := helper.APIResponseWithoutData(http.StatusNotFound, false, "Data tidak ditemukan")
		c.JSON(http.StatusNotFound, response)
		return
	}

	responseData := shareddomain.ToResponseStudentRegistrationByNimPeriodID(registrationStudent)
	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Riwayat pendaftaran berhasil didapatkan", responseData)
	c.JSON(http.StatusOK, response)
}

func (handler *StudentRegistrationResthandlerImpl) FindStudentRegistrationByPeriod(c *gin.Context) {
	// get period from param
	periodID := c.Param("periodID")

	// call service
	studentRegistration, err := handler.service.FindStudentRegistrationByPeriod(periodID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Pendaftaran mahasiswa gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responseData := []shareddomain.StudentRegistrationPeriodResponse{}
	for _, data := range studentRegistration {
		responseData = append(responseData, shareddomain.ToStudentRegistrationPeriod(data))
	}
	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Riwayat pendaftaran berhasil didapatkan", responseData)
	c.JSON(http.StatusOK, response)
}

func (handler *StudentRegistrationResthandlerImpl) UpdateStudentRegistration(c *gin.Context) {
	//get id in param
	var request shareddomain.UpdateStudentRegistrationRequest
	request.ID = c.Param("ID")

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Validasi gagal dilakukan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.UpdateStudentRegistration(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Validasi gagal dilakukan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "Validasi berhasil dilakukan")
	c.JSON(http.StatusCreated, response)
}
