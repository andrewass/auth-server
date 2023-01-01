package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpClientRoutes(router *gin.Engine) {
	router.GET("/clients", getClientsHandler)
	router.POST("/clients", createClientHandler)
	router.PATCH("/clients", updateClientHandler)
	router.DELETE("/clients", deleteClientHandler)
}

const includeSecret = true
const notIncludeSecret = false

func getClientsHandler(context *gin.Context) {
	email := context.Query("email")
	var clients = getClients(email)
	mappedClients := make([]ClientInformationDto, len(clients))
	for i, client := range clients {
		mappedClients[i] = mapToClientResponse(client, notIncludeSecret)
	}
	context.JSON(http.StatusOK, mappedClients)
}

func createClientHandler(context *gin.Context) {
	request := AddClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	request.ClientSecret = generateRandomString()
	client := createNewClient(request)
	context.JSON(http.StatusOK, mapToClientResponse(client, includeSecret))
}

func updateClientHandler(context *gin.Context) {
	request := UpdateClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := updateClient(request)
	context.JSON(http.StatusOK, mapToClientResponse(client, notIncludeSecret))
}

func deleteClientHandler(context *gin.Context) {
	clientID := context.Query("client_id")
	deleteClient(clientID)
	context.Status(http.StatusOK)
}

func mapToClientResponse(client Client, includeSecret bool) ClientInformationDto {
	clientDto := ClientInformationDto{
		ClientId:                client.ClientId,
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
		UserEmail:               client.UserEmail,
	}
	if includeSecret {
		clientDto.ClientSecret = client.ClientSecret
	}
	return clientDto
}
