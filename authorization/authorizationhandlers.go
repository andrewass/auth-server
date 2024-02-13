package authorization

import (
	"auth-server/authorization/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthorizationHandler struct {
	Service *AuthorizationService
}

func (h *AuthorizationHandler) SetUpAuthorizationRoutes(router *gin.Engine) {
	router.POST("/server/auth/response", h.authorizationCodeResponseHandler)
}

func (h *AuthorizationHandler) authorizationCodeResponseHandler(context *gin.Context) {
	var request dto.AuthorizationCodeRequest
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	response := h.Service.createAndSaveAuthorizationCode(request)
	context.JSON(http.StatusOK, response)
}
