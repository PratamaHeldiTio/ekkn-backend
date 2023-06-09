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

func (handler *lecturerRegistrationRestHandlerImpl) FindLecturerRegistrationHistory(c *gin.Context) {
	lecturerID := c.MustGet("lecturerID").(string)

	lecturerRegistrations, err := handler.service.FindLecturerRegistrationByLectureID(lecturerID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "riwayat pendaftaran gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	lecturerRegistrationResponses := []shareddomain.LecturerRegistrationHistoryResponse{}
	for _, lecturerRegistration := range lecturerRegistrations {
		lecturerRegistrationResponses = append(lecturerRegistrationResponses, shareddomain.ToLecturerRegistrationHistory(lecturerRegistration))
	}
	response := helper.APIResponseWithData(http.StatusOK, true, "riwayat pendaftaran berhasil didapatkan", lecturerRegistrationResponses)
	c.JSON(http.StatusOK, response)

}

func (handler *lecturerRegistrationRestHandlerImpl) ValidationLecturerRegistration(c *gin.Context) {
	var request shareddomain.ValidationLectureRegistrationRequest
	request.ID = c.Param("id")

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "validasi gagal dilakukan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.ValidationLecturerRegistration(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "validasi gagal dilakukan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "validasi berhasil dilakukan")
	c.JSON(http.StatusCreated, response)
}

func (handler *lecturerRegistrationRestHandlerImpl) FindLecturerRegistrationByPeriod(c *gin.Context) {
	PeriodID := c.Param("periodID")
	query := c.Query("search")

	lecturerRegistrations, err := handler.service.FindLecturerRegistrationByPeriod(PeriodID, query)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "pendaftaran dosen gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	lecturerRegistrationResponses := []shareddomain.LecturerRegistrationByPeriodResponse{}
	for _, lecturerRegistration := range lecturerRegistrations {
		lecturerRegistrationResponses = append(lecturerRegistrationResponses, shareddomain.ToLecturerRegistrationByPeriod(lecturerRegistration))
	}
	response := helper.APIResponseWithData(http.StatusOK, true, "pendaftaran dosen berhasil didapatkan", lecturerRegistrationResponses)
	c.JSON(http.StatusOK, response)
}

func (handler *lecturerRegistrationRestHandlerImpl) FindLecturerRegistrationByPeriodApprove(c *gin.Context) {
	PeriodID := c.Param("periodID")

	lecturerRegistrations, err := handler.service.FindLecturerRegistrationByPeriod(PeriodID, "")
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "dosen gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	lecturerRegistrationResponses := []shareddomain.LecturerRegistrationByPeriodResponse{}
	for _, lecturerRegistration := range lecturerRegistrations {
		if lecturerRegistration.Status == "true" {
			lecturerRegistrationResponses = append(lecturerRegistrationResponses, shareddomain.ToLecturerRegistrationByPeriod(lecturerRegistration))
		}
	}
	response := helper.APIResponseWithData(http.StatusOK, true, "dosen berhasil didapatkan", lecturerRegistrationResponses)
	c.JSON(http.StatusOK, response)
}

func (handler *lecturerRegistrationRestHandlerImpl) FindLecturerRegistrationApprove(c *gin.Context) {
	lecturerID := c.MustGet("lecturerID").(string)

	lecturerRegistrations, err := handler.service.FindLecturerRegistrationByLectureID(lecturerID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "riwayat pendaftaran gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// response valid
	lecturerRegistrationResponses := []shareddomain.LecturerRegistrationApproveResponse{}
	for _, lecturerRegistration := range lecturerRegistrations {
		if lecturerRegistration.Status == "true" {
			lecturerRegistrationResponses = append(lecturerRegistrationResponses, shareddomain.ToLecturerRegistrationApprove(lecturerRegistration))
		}
	}
	response := helper.APIResponseWithData(http.StatusOK, true, "riwayat pendaftaran berhasil didapatkan", lecturerRegistrationResponses)
	c.JSON(http.StatusOK, response)
}
