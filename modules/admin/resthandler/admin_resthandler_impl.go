package resthandler

import (
	jwtmanager "backend-ekkn/jwt_manager"
	"backend-ekkn/modules/admin/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminResthandlerImpl struct {
	service service.AdminService
	jwt     jwtmanager.JwtManager
}

func NewAdminRestHandler(service service.AdminService, jwt jwtmanager.JwtManager) AdminRestHandler {
	return &AdminResthandlerImpl{service, jwt}
}

func (handler *AdminResthandlerImpl) CreateAdmin(c *gin.Context) {
	var request shareddomain.AdminRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Admin gagal ditambahkan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.CreateAdmin(request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Admin gagal ditambahkan", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusCreated, true, "Admin berhasil ditambahkan")
	c.JSON(http.StatusCreated, response)
}

func (handler *AdminResthandlerImpl) LoginAdmin(c *gin.Context) {
	var request shareddomain.AdminRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusUnprocessableEntity, false, "Admin gagal ditambahkan", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	admin, err := handler.service.LoginAdmin(request)
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Data yang anda masukan salah", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := handler.jwt.GenerateJwt(admin.Username, "admin")
	if err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusInternalServerError, false, "Gagal login silahkan coba lagi", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//access jwt_manager for response data
	accessToken := gin.H{"access_token": token}
	// create response
	response := helper.APIResponseWithData(http.StatusOK, true, "Admin berhasil login", accessToken)

	c.JSON(http.StatusOK, response)
}

func (handler *AdminResthandlerImpl) DeleteAdmin(c *gin.Context) {
	username := c.Param("username")

	if err := handler.service.DeleteAdmin(username); err != nil {
		// create response
		response := helper.APIResponseWithError(http.StatusBadRequest, false, "Admin gagal dihapus", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create response
	response := helper.APIResponseWithoutData(http.StatusOK, true, "Admin berhasil dihapus")
	c.JSON(http.StatusOK, response)
}
