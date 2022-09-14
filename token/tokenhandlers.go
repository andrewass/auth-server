package token

import (
	"auth-server/token/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpTokenRoutes(router *gin.Engine) {
	router.POST("/token", getTokenHandler)
	router.POST("/token/introspect", introspectTokenHandler)
	router.POST("token/revoke", revokeTokenHandler)
}

func getTokenHandler(context *gin.Context) {
	var request dto.GetTokenRequest
	err := context.Bind(&request)
	if err != nil {
		panic(err)
	}
	response := getToken(request)
	context.JSON(http.StatusOK, response)
}

func introspectTokenHandler(context *gin.Context) {
	request := dto.IntrospectTokenRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	response := introspectToken(request)
	context.JSON(http.StatusOK, response)
}

func revokeTokenHandler(context *gin.Context) {
	request := dto.RevokeTokenRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	revokeToken(request)
	context.Status(http.StatusOK)
}
