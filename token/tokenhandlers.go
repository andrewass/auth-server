package token

import (
	"auth-server/token/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpTokenRoutes(router *gin.Engine) {
	router.POST("/token", getTokensHandler)
	router.POST("/token/introspect", introspectTokenHandler)
	router.POST("token/revoke", revokeTokenHandler)
}

func getTokensHandler(context *gin.Context) {
	var request types.GetTokenRequest
	err := context.Bind(&request)
	if err != nil {
		panic(err)
	}
	response := getTokens(request)
	context.JSON(http.StatusOK, response)
}

func introspectTokenHandler(context *gin.Context) {
	request := types.IntrospectTokenRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	response := introspectToken(request)
	context.JSON(http.StatusOK, response)
}

func revokeTokenHandler(context *gin.Context) {
	request := types.RevokeTokenRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	revokeToken(request)
	context.Status(http.StatusOK)
}
