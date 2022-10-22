package authorization

import (
	"auth-server/authorization/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpAuthorizationRoutes(router *gin.Engine) {
	router.GET("/authorize", authorizeUserHandler)
	router.POST("/authorization-response", authorizationCodeResponseHandler)
}

func authorizeUserHandler(context *gin.Context) {
	var request dto.AuthorizeRequest
	err := context.BindQuery(&request)
	if err != nil {
		panic(err)
	}
	frontendUrl := createFrontendUrl(request)
	context.Redirect(http.StatusFound, frontendUrl)
}

func authorizationCodeResponseHandler(context *gin.Context) {
	var request dto.AuthorizationCodeRequest
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	authorizationCode := createAndSaveAuthorizationCode(request.Email, request.ClientId)
	context.JSON(http.StatusOK, authorizationCode)
}
