package token

import (
	"auth-server/token/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenHandlers struct {
	Service *TokenService
}

func (h *TokenHandlers) SetUpTokenRoutes(router *gin.Engine) {
	router.POST("/api/token/token", h.getTokensHandler)
	router.POST("/token/introspect", h.introspectTokenHandler)
	router.POST("token/revoke", h.revokeTokenHandler)
}

func (h *TokenHandlers) getTokensHandler(context *gin.Context) {
	var request types.GetTokenRequest
	err := context.Bind(&request)
	if err != nil {
		panic(err)
	}
	response := h.Service.processGetTokensRequest(request)
	context.JSON(http.StatusOK, response)
}

func (h *TokenHandlers) introspectTokenHandler(context *gin.Context) {
	request := types.IntrospectTokenRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	response := h.Service.introspectToken(request)
	context.JSON(http.StatusOK, response)
}

func (h *TokenHandlers) revokeTokenHandler(context *gin.Context) {
	request := types.RevokeTokenRequest{}
	err := context.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	h.Service.revokeToken(request)
	context.Status(http.StatusOK)
}
