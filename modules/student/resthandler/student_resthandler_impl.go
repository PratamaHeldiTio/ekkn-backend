package resthandler

import (
	"backend-ekkn/jwt_manager"
	"backend-ekkn/modules/student/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudentResthandlerImpl struct {
	service     service.StudentService
	authService jwtmanager.JwtManager
}

func NewStudentResthandler(service service.StudentService, authService jwtmanager.JwtManager) StudentResthandler {
	return &StudentResthandlerImpl{
		service,
		authService,
	}
}

func (handler *StudentResthandlerImpl) CreateStudent(c *gin.Context) {
	var studentRequest shareddomain.CreateStudent

	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusUnprocessableEntity, false, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send data to service and get return
	student, err := handler.service.CreateStudent(studentRequest)
	if err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusBadRequest, false, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//map domain to respon data
	responseStudent := shareddomain.ToResponseStudent(student)

	// create response
	response := helper.APIResponseSuccess(http.StatusCreated, true, responseStudent)

	c.JSON(http.StatusCreated, response)
}

func (handler *StudentResthandlerImpl) FindStudentByNim(c *gin.Context) {
	// get params from url path
	nim := c.Param("nim")

	// send data and get return from service
	student, err := handler.service.FindStudentByNim(nim)
	if err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusNotFound, false, err.Error())

		c.JSON(http.StatusNotFound, response)
		return
	}

	// map data domain to response data
	responseStudent := shareddomain.ToResponseFindStudentByNim(student)

	// create response
	response := helper.APIResponseSuccess(http.StatusOK, true, responseStudent)

	c.JSON(http.StatusOK, response)
}

func (handler *StudentResthandlerImpl) LoginStudent(c *gin.Context) {
	var studentRequest shareddomain.LoginStudent

	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusUnprocessableEntity, false, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send data and get return from service
	student, err := handler.service.LoginStudent(studentRequest)
	if err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusBadRequest, false, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// jwt service generation
	token, err := handler.authService.GenerateJwt(student.Nim, "student")
	if err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusBadRequest, false, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//access jwt_manager for response data
	accessToken := gin.H{"access_token": token}
	// create response
	response := helper.APIResponseSuccess(http.StatusOK, true, accessToken)

	c.JSON(http.StatusOK, response)
}

func (handler *StudentResthandlerImpl) FindAllStudent(c *gin.Context) {
	// send data to service and get return
	students, err := handler.service.FindAllStudent()
	if err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusInternalServerError, false, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// map data slice to response student slice
	responseStudent := []shareddomain.CreateStudent{}
	for _, student := range students {
		responseStudent = append(responseStudent, shareddomain.ToResponseStudent(student))
	}

	response := helper.APIResponseSuccess(http.StatusOK, true, responseStudent)
	c.JSON(http.StatusOK, response)
}

func (handler *StudentResthandlerImpl) UpdateStudent(c *gin.Context) {
	var studentRequest shareddomain.UpdateStudent

	// get id from url
	nim := c.Param("nim")

	if err := c.ShouldBindJSON(&studentRequest); err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusUnprocessableEntity, false, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// asign nim to struct
	studentRequest.Nim = nim

	student, err := handler.service.UpdateStudent(studentRequest)

	if err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusNotFound, false, err.Error())
		c.JSON(http.StatusNotFound, response)
		return
	}
	responseStudent := shareddomain.ToResponseUpdateStudent(student)

	// create response
	response := helper.APIResponseSuccess(http.StatusOK, true, responseStudent)

	c.JSON(http.StatusOK, response)

}

func (handler *StudentResthandlerImpl) DeleteStudent(c *gin.Context) {
	// get id
	nim := c.Param("nim")

	if err := handler.service.DeleteStudent(nim); err != nil {
		// create response
		response := helper.APIResponseFail(http.StatusBadRequest, false, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseSuccess(http.StatusOK, true, nil)

	c.JSON(http.StatusOK, response)
}
