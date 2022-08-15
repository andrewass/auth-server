package client

import (
	"auth-server/client/dto"
	"auth-server/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetUpClientRoutes(router *gin.Engine) {
	//Get all registered OAuth 2.0 clients / OpenID relying parties.
	router.GET("/clients", getClientsHandler)
	//Registers a new OAuth 2.0 client / OpenID relying on party.
	router.POST("/clients", addClientHandler)
}

func getClientsHandler(ctx *gin.Context) {
	validateMasterAccessToken(ctx)
	var client = getClients()
	ctx.JSON(http.StatusOK, client)
}

func addClientHandler(ctx *gin.Context) {
	validateMasterAccessToken(ctx)
	request := dto.AddClientRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		panic(err)
	}
	addClient(request)
	ctx.Status(http.StatusOK)
}

func validateMasterAccessToken(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")
	if !common.ValidateMasterAccessToken(token[1]) {
		ctx.Status(http.StatusUnauthorized)
		return
	}
}
