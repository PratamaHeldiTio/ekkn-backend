package resthandler

import (
	"backend-ekkn/modules/logbook/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogbookRestHandlerImpl struct {
	service service.LogbookService
}

func NewLogbookResthandler(service service.LogbookService) LogbookRestHandler {
	return &LogbookRestHandlerImpl{service}
}

func (handler *LogbookRestHandlerImpl) CreateLogbook(c *gin.Context) {
	// bind request
	var request shareddomain.LogbookRequest
	//if err := c.ShouldBindJSON(&request); err != nil {
	//	// create response
	//	response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Logbook gagal ditambahkan", err.Error())
	//	c.JSON(http.StatusUnprocessableEntity, response)
	//	return
	//}

	// call service create logbook
	if err := handler.service.CreateLogbook(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Logbook gagal ditambahkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// valid
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "Logbook gagal ditambahkan")
	c.JSON(http.StatusCreated, response)
}

func (handler *LogbookRestHandlerImpl) FindLogbookByStudentPeriod(c *gin.Context) {
	studentID := c.MustGet("currentUser").(string)
	periodID := c.Param("periodID")

	// call service
	logbooks, err := handler.service.FindLogbookByStudentPeriod(studentID, periodID)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Logbook gagal didapatkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	logbookResponses := []shareddomain.LogbookResponse{}
	for _, logbook := range logbooks {
		logbookResponses = append(logbookResponses, shareddomain.ToLogbookResponse(logbook))
	}

	//reponse
	response := helper.APIResponseWithData(http.StatusOK, true, "Logbook berhasil didapatkan", logbookResponses)
	c.JSON(http.StatusOK, response)
}
