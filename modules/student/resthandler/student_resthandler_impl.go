package resthandler

import (
	"backend-ekkn/modules/student/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentResthandlerImpl struct {
	service service.StudentService
}

func NewStudentResthandler(service service.StudentService) StudentResthandler {
	return &StudentResthandlerImpl{service}
}

func (handler *StudentResthandlerImpl) CreateStudent(c *gin.Context) {
	var studentRequest shareddomain.CreateStudentRequest

	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		errors := helper.FormatValidationError(err)
		errorData := gin.H{"errors": errors}
		response := helper.APIResponse(http.StatusUnprocessableEntity, false, "Mahasiswa gagal ditambahkan", errorData)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	student, err := handler.service.CreateStudent(studentRequest)
	if err != nil {
		errorData := gin.H{"errors": err.Error()}
		response := helper.APIResponse(http.StatusBadRequest, false, "Mahasiswa gagal ditambahkan", errorData)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// jwt service generation

	//map domain to respon data
	responseStudent := shareddomain.ToResponseStudent(student, "sfsffdfds")
	// build api respon
	response := helper.APIResponse(http.StatusCreated, true, "Mahasiswa berhasil ditambahkan", responseStudent)

	c.JSON(http.StatusOK, response)
}
