package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpClientRoutes(router *gin.Engine) {
	router.GET("/client", getClientHandler)
	router.GET("/clients", getAllClientsHandler)
	router.POST("/add-client", addClientHandler)

}

func getClientHandler(context *gin.Context) {
	var clientId = context.Query("clientId")
	var client = GetClient(clientId)
	context.JSON(http.StatusOK, client)
}

func getAllClientsHandler(context *gin.Context) {

}

func addClientHandler(context *gin.Context) {

}
