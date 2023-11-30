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
	router.POST("/api/auth/response", h.authorizationCodeResponseHandler)
}

func (h *AuthorizationHandler) authorizeUserHandler(context *gin.Context) {
	var request dto.AuthorizeRequest
	err := context.BindQuery(&request)
	if err != nil {
		panic(err)
	}
	frontendUrl := h.Service.createFrontendUrl(request)
	context.Redirect(http.StatusFound, frontendUrl)
}

func (h *AuthorizationHandler) authorizationCodeResponseHandler(context *gin.Context) {
	var request dto.AuthorizationCodeRequest
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	authorizationCode := h.Service.createAndSaveAuthorizationCode(request)
	context.JSON(http.StatusOK, gin.H{"code": authorizationCode})
}
