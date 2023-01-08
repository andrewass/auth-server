package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClientHandlers struct {
	Service *ClientService
}

func (h *ClientHandlers) SetUpClientRoutes(router *gin.Engine) {
	router.GET("/clients", h.getClientsHandler)
	router.POST("/clients", h.createClientHandler)
	router.PATCH("/clients", h.updateClientHandler)
	router.DELETE("/clients", h.deleteClientHandler)
	router.POST("/client/rotate-secret", h.rotateClientSecretHandler)
}

const includeSecret = true
const notIncludeSecret = false

func (h *ClientHandlers) getClientsHandler(context *gin.Context) {
	email := context.Query("email")
	var clients = h.Service.getClients(email)
	mappedClients := make([]ClientInformationDto, len(clients))
	for i, client := range clients {
		mappedClients[i] = h.mapToClientResponse(client, notIncludeSecret)
	}
	context.JSON(http.StatusOK, mappedClients)
}

func (h *ClientHandlers) createClientHandler(context *gin.Context) {
	request := AddClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	request.ClientSecret = generateRandomString()
	client := h.Service.createNewClient(request)
	context.JSON(http.StatusOK, h.mapToClientResponse(client, includeSecret))
}

func (h *ClientHandlers) updateClientHandler(context *gin.Context) {
	request := UpdateClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := h.Service.updateClient(request)
	context.JSON(http.StatusOK, h.mapToClientResponse(client, notIncludeSecret))
}

func (h *ClientHandlers) deleteClientHandler(context *gin.Context) {
	clientId := context.Query("client_id")
	h.Service.deleteClient(clientId)
	context.Status(http.StatusOK)
}

func (h *ClientHandlers) rotateClientSecretHandler(context *gin.Context) {
	clientId := context.Query("client_id")
	client := h.Service.rotateClientSecret(clientId)
	context.JSON(http.StatusOK, h.mapToClientResponse(client, includeSecret))
}

func (h *ClientHandlers) AddAdminClient() {
	h.Service.AddAdminClient()
}

func (h *ClientHandlers) mapToClientResponse(client Client, includeSecret bool) ClientInformationDto {
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
