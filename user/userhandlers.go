package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpUserRoutes(router *gin.Engine) {
	router.POST("/user/sign-up", getOpenIdConfigHandler)
}

func getOpenIdConfigHandler(context *gin.Context) {
	response := "fdf"
	context.JSON(http.StatusOK, response)
}

func getOauthConfigHandler(context *gin.Context) {
	response := "fdf"
	context.JSON(http.StatusOK, response)
}
