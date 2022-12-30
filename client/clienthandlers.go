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

func getClientsHandler(context *gin.Context) {
	email := context.Query("email")
	var clients = getClients(email)
	mappedClients := make([]ClientResponse, len(clients))
	for i, client := range clients {
		mappedClients[i] = mapToClientResponse(client)
	}
	context.JSON(http.StatusOK, mappedClients)
}

func createClientHandler(context *gin.Context) {
	request := AddClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := createNewClient(request)
	context.JSON(http.StatusOK, mapToClientResponse(client))
}

func updateClientHandler(context *gin.Context) {
	request := UpdateClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := updateClient(request)
	context.JSON(http.StatusOK, mapToClientResponse(client))
}

func deleteClientHandler(context *gin.Context) {
	request := DeleteClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	deleteClient(request.ClientID, request.ClientSecret)
	context.Status(http.StatusOK)
}

func mapToClientResponse(client Client) ClientResponse {
	return ClientResponse{
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
		UserEmail:               client.UserEmail,
	}
}
