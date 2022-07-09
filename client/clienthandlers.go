package client

import (
	"auth-server/client/dto"
	"auth-server/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetUpClientRoutes(router *gin.Engine) {
	//Get a previous registered OAuth 2.0 client / OpenID relying party.
	router.GET("/clients", getClientHandler)
	//Registers a new OAuth 2.0 client / OpenID relying party.
	router.POST("/clients", addClientHandler)
}

func getClientHandler(context *gin.Context) {
	var clientId = context.Query("clientId")
	var client = getClient(clientId)
	context.JSON(http.StatusOK, client)
}

func addClientHandler(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")
	if !common.ValidateMasterAccessToken(token[1]) {
		context.Status(http.StatusUnauthorized)
		return
	}
	request := dto.AddClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	addClient(request)
	context.Status(http.StatusOK)
}
