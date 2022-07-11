package discovery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpDiscoveryRoutes(router *gin.Engine) {
	//Discover the OpenID Connect endpoints, capabilities, supported cryptographic algorithms and features.
	router.GET("/.well-known/openid-configuration", getOpenIdConfigHandler)
	//Discover the OAuth 2.0 endpoints, capabilities, supported cryptographic algorithms and features.
	router.GET("/.well-known/oauth-authorization-server", getOauthConfigHandler)
}

func getOpenIdConfigHandler(context *gin.Context) {
	response := getOpenIdConfig()
	context.JSON(http.StatusOK, response)
}

func getOauthConfigHandler(context *gin.Context) {
	response := getOauthConfig()
	context.JSON(http.StatusOK, response)
}
