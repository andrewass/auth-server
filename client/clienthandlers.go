package client

import (
	"auth-server/client/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpClientRoutes(router *gin.Engine) {
	router.GET("/clients", getClientsHandler)
	router.POST("/clients", addClientHandler)
	router.PATCH("/clients", updateClientHandler)
}

func getClientsHandler(context *gin.Context) {
	email := context.Query("email")
	var clients = getClients(email)
	mappedClients := make([]dto.ClientResponse, len(clients))
	for i, client := range clients {
		mappedClients[i] = mapToClientResponse(client)
	}
	context.JSON(http.StatusOK, mappedClients)
}

func addClientHandler(context *gin.Context) {
	request := dto.AddClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := addClient(request)
	context.JSON(http.StatusOK, mapToClientResponse(client))
}

func updateClientHandler(context *gin.Context) {
	request := dto.UpdateClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := updateClient(request)
	context.JSON(http.StatusOK, mapToClientResponse(client))
}

func mapToClientResponse(client Client) dto.ClientResponse {
	return dto.ClientResponse{
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
