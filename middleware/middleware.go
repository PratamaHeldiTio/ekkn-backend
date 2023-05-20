package middleware

import (
	jwtmanager "backend-ekkn/jwt_manager"
	"backend-ekkn/pkg/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	JwtManager jwtmanager.JwtManager
}

func NewAtuhMiddleware(jwtManager jwtmanager.JwtManager) *AuthMiddleware {
	return &AuthMiddleware{jwtManager}
}

func (auth *AuthMiddleware) AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		authHeader := c.GetHeader("Authorization")

		// check bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get token authheader "Bearer token"
		var tokenStr string
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenStr = tokenArray[1]
		}

		// validate
		token, err := auth.JwtManager.ValidateJwt(tokenStr)
		if err != nil {
			response := helper.APIResponseWithError(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get payload from jwt
		claim, ok := token.Claims.(jwt.MapClaims)

		// check validation token and get payload token
		if !token.Valid || !ok {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// mapping payload id to var
		userID := claim["id"].(string)
		userRole := claim["role"].(string)

		// set with context gin
		c.Set("currentUser", userID)

		//role
		c.Set("currentRole", userRole)
	}
}

func (auth *AuthMiddleware) AuthMiddleWareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		authHeader := c.GetHeader("Authorization")

		// check bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get token authheader "Bearer token"
		var tokenStr string
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenStr = tokenArray[1]
		}

		// validate
		token, err := auth.JwtManager.ValidateJwt(tokenStr)
		if err != nil {
			response := helper.APIResponseWithError(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get payload from jwt
		claim, ok := token.Claims.(jwt.MapClaims)

		// check validation token and get payload token
		if !token.Valid || !ok {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// mapping payload id to var
		userRole := claim["role"].(string)
		if userRole != "admin" {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

	}
}

func (auth *AuthMiddleware) AuthMiddleWareLecturer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		authHeader := c.GetHeader("Authorization")

		// check bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get token authheader "Bearer token"
		var tokenStr string
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenStr = tokenArray[1]
		}

		// validate
		token, err := auth.JwtManager.ValidateJwt(tokenStr)
		if err != nil {
			response := helper.APIResponseWithError(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get payload from jwt
		claim, ok := token.Claims.(jwt.MapClaims)

		// check validation token and get payload token
		if !token.Valid || !ok {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// mapping payload id to var
		userRole := claim["role"].(string)
		userID := claim["id"].(string)
		if userRole != "lecturer" {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// set with context gin
		c.Set("lecturerID", userID)
	}
}

func (auth *AuthMiddleware) AuthMiddleWareLecturerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get authorization header
		authHeader := c.GetHeader("Authorization")

		// check bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get token authheader "Bearer token"
		var tokenStr string
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenStr = tokenArray[1]
		}

		// validate
		token, err := auth.JwtManager.ValidateJwt(tokenStr)
		if err != nil {
			response := helper.APIResponseWithError(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// get payload from jwt
		claim, ok := token.Claims.(jwt.MapClaims)

		// check validation token and get payload token
		if !token.Valid || !ok {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// mapping payload id to var
		userRole := claim["role"].(string)
		if userRole == "student" {
			response := helper.APIResponseWithoutData(http.StatusUnauthorized, false, "Anda tidak memiliki akses ini")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
