package token

import (
	"auth-server/token/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpTokenRoutes(router *gin.Engine) {
	router.POST("/token/introspect", introspectTokenHandler)
	router.POST("token/revoke", revokeTokenHandler)
}

func introspectTokenHandler(context *gin.Context) {
	request := dto.IntrospectTokenRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	var response = introspectToken(request)
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
