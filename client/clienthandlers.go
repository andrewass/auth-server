package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClientHandlers struct {
	Service *ClientService
}

func (h *ClientHandlers) SetUpClientRoutes(router *gin.Engine) {
	router.GET("/server/clients", h.getClientsHandler)
	router.GET("/server/clients/client", h.getClientHandler)
	router.POST("/server/clients/create", h.createClientHandler)
	router.PATCH("server/clients", h.updateClientHandler)
	router.DELETE("server/clients", h.deleteClientHandler)
	router.POST("server/clients/rotate-secret", h.rotateClientSecretHandler)
}

func (h *ClientHandlers) getClientsHandler(context *gin.Context) {
	email := context.Query("email")
	var clients = h.Service.getClients(email)
	mappedClients := make([]ClientInformationDto, len(clients))
	for i, client := range clients {
		mappedClients[i] = h.mapToClientResponse(client)
	}
	context.JSON(http.StatusOK, mappedClients)
}

func (h *ClientHandlers) getClientHandler(context *gin.Context) {
	clientId := context.Query("client_id")
	var client = h.Service.GetClientById(clientId)
	context.JSON(http.StatusOK, h.mapToClientResponse(client))
}

func (h *ClientHandlers) createClientHandler(context *gin.Context) {
	request := AddClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	request.ClientSecret = generateRandomString()
	client := h.Service.createNewClient(request)
	context.JSON(http.StatusOK, h.mapToClientResponse(client))
}

func (h *ClientHandlers) updateClientHandler(context *gin.Context) {
	request := UpdateClientRequest{}
	if err := context.BindJSON(&request); err != nil {
		panic(err)
	}
	client := h.Service.updateClient(request)
	context.JSON(http.StatusOK, h.mapToClientResponse(client))
}

func (h *ClientHandlers) deleteClientHandler(context *gin.Context) {
	clientId := context.Query("client_id")
	h.Service.deleteClient(clientId)
	context.Status(http.StatusOK)
}

func (h *ClientHandlers) rotateClientSecretHandler(context *gin.Context) {
	clientId := context.Query("client_id")
	client := h.Service.rotateClientSecret(clientId)
	context.JSON(http.StatusOK, h.mapToClientResponse(client))
}

func (h *ClientHandlers) AddAdminClient() {
	h.Service.AddAdminClient()
}

func (h *ClientHandlers) mapToClientResponse(client Client) ClientInformationDto {
	clientDto := ClientInformationDto{
		ClientId:                client.ClientId,
		ClientSecret:            client.ClientSecret,
		ClientIdIssuedAt:        client.ClientIdIssuedAt,
		ClientSecretIssuedAt:    client.ClientSecretIssuedAt,
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
	return clientDto
}
