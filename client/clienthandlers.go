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

func getClientsHandler(context *gin.Context) {
	validateMasterAccessToken(context)
	var client = getClients()
	context.JSON(http.StatusOK, client)
}

func addClientHandler(context *gin.Context) {
	validateMasterAccessToken(context)
	request := dto.AddClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := addClient(request)
	context.JSON(http.StatusOK, mapToAddClientResponse(client))
}

func validateMasterAccessToken(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")
	if !common.ValidateMasterAccessToken(token[1]) {
		context.Status(http.StatusUnauthorized)
		return
	}
}

func mapToAddClientResponse(client Client) dto.AddClientResponse {
	return dto.AddClientResponse{
		ClientId:                client.ClientId,
		ClientSecret:            client.ClientSecret,
		ClientIdIssuedAt:        client.ClientIdIssuedAt,
		LogoUri:                 client.LogoUri,
		ApplicationType:         client.ApplicationType,
		ClientName:              client.ClientName,
		ClientUri:               client.ClientUri,
		InitiateLoginUri:        client.InitiateLoginUri,
		TokenEndpointAuthMethod: client.TokenEndpointAuthMethod,
		RedirectUris:            client.RedirectUris,
		ResponseTypes:           client.ResponseTypes,
		GrantTypes:              client.GrantTypes,
		PostLogoutRedirectUris:  client.PostLogoutRedirectUris,
	}
}
