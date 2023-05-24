package resthandler

import (
	jwtmanager "backend-ekkn/jwt_manager"
	"backend-ekkn/modules/lecturer/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LecturerRestHandlerImpl struct {
	service     service.LecturerService
	authService jwtmanager.JwtManager
}

func NewLecturerRestHandler(service service.LecturerService, authService jwtmanager.JwtManager) LecturerRestHandler {
	return &LecturerRestHandlerImpl{service, authService}
}

func (handler *LecturerRestHandlerImpl) CreateLecturer(c *gin.Context) {
	// get request
	var request shareddomain.LecturerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "dosen gagal ditambahkan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.CreateLecturer(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "dosen gagal ditambahkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "dosen berhasil ditambahkan")
	c.JSON(http.StatusCreated, response)
	return
}

func (handler *LecturerRestHandlerImpl) UpdateLecturerByJwt(c *gin.Context) {
	var request shareddomain.LecturerRequest
	request.ID = c.MustGet("lecturerID").(string)

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "dosen gagal diedit", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.UpdateLecturer(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "dosen gagal diedit", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "dosen berhasil diedit")
	c.JSON(http.StatusOK, response)
	return
}

func (handler *LecturerRestHandlerImpl) UpdateLecturer(c *gin.Context) {
	var request shareddomain.LecturerRequest
	request.ID = c.Param("id")

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "dosen gagal diedit", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.UpdateLecturer(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "dosen gagal diedit", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "dosen berhasil diedit")
	c.JSON(http.StatusOK, response)
	return
}

func (handler *LecturerRestHandlerImpl) DeleteLecturer(c *gin.Context) {
	ID := c.Param("id")

	if err := handler.service.DeleteLecture(ID); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "dosen gagal dihapus", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "dosen berhasil dihapus")
	c.JSON(http.StatusOK, response)
	return
}

func (handler *LecturerRestHandlerImpl) ResetPassword(c *gin.Context) {
	ID := c.Param("id")

	if err := handler.service.ResetPassword(ID); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal reset password", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil reset password")
	c.JSON(http.StatusOK, response)
}

func (handler *LecturerRestHandlerImpl) FindByIdParam(c *gin.Context) {
	ID := c.Param("id")

	lecturer, err := handler.service.FindLecturerByID(ID)
	if err != nil {
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mendapatkan dosen", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	lecturerResponse := shareddomain.ToLecturerResponse(lecturer)
	response := helper.APIResponseWithData(http.StatusOK, true, "berhasil mendapatkan dosen", lecturerResponse)
	c.JSON(http.StatusOK, response)

}

func (handler *LecturerRestHandlerImpl) FindAllLecturer(c *gin.Context) {
	query := c.Query("search")
	lecturers, err := handler.service.FindAllLecturer(query)
	if err != nil {
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mendapatkan dosen", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	lecturersResponse := []shareddomain.LecturerResponse{}
	for _, lecturer := range lecturers {
		lecturersResponse = append(lecturersResponse, shareddomain.ToLecturerResponse(lecturer))
	}
	response := helper.APIResponseWithData(http.StatusOK, true, "berhasil mendapatkan dosen", lecturersResponse)
	c.JSON(http.StatusOK, response)
}

func (handler *LecturerRestHandlerImpl) LoginLecturer(c *gin.Context) {
	var request shareddomain.LecturerLogin

	// validation with gin validator playground golang/v10
	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Data yang anda masukan salah", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send data and get return from service
	lecturer, err := handler.service.LoginLecturer(request)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Data yang anda masukan salah", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// jwt service generation
	token, err := handler.authService.GenerateJwt(lecturer.ID, "lecturer")
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusInternalServerError, false, "Gagal login silahkan coba lagi", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//access jwt_manager for response data
	accessToken := gin.H{"access_token": token}
	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Dosen berhasil login", accessToken)
	c.JSON(http.StatusOK, response)
}

func (handler *LecturerRestHandlerImpl) FindByIdJwt(c *gin.Context) {
	ID := c.MustGet("lecturerID").(string)

	lecturer, err := handler.service.FindLecturerByID(ID)
	if err != nil {
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal mendapatkan dosen", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	lecturerResponse := shareddomain.ToLecturerResponse(lecturer)
	response := helper.APIResponseWithData(http.StatusOK, true, "berhasil mendapatkan dosen", lecturerResponse)
	c.JSON(http.StatusOK, response)

}

func (handler *LecturerRestHandlerImpl) ChangePassword(c *gin.Context) {
	var request shareddomain.ChangePasswordLecturerRequest
	request.ID = c.MustGet("lecturerID").(string)

	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Data yang anda masukan salah", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.ChangePassword(request); err != nil {
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "gagal merubah password", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "berhasil merubah password")
	c.JSON(http.StatusOK, response)
}
