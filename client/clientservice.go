package client

import (
	"auth-server/client/dto"
	"math/rand"
	"time"
)

var clients = []Client{
	{ClientId: "firstId", ClientSecret: "firstSecret"},
	{ClientId: "secId", ClientSecret: "secSecret"},
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getClients() []Client {
	return getAllClients()
}

func addClient(request dto.AddClientRequest) Client {
	client := Client{
		ClientId:                generateRandomString(),
		ClientSecret:            generateRandomString(),
		ClientIdIssuedAt:        time.Now().Unix(),
		LogoUri:                 request.LogoUri,
		ApplicationType:         request.ApplicationType,
		ClientName:              request.ClientName,
		ClientUri:               request.ClientUri,
		InitiateLoginUri:        request.InitiateLoginUri,
		TokenEndpointAuthMethod: request.TokenEndpointAuthMethod,
		RedirectUris:            request.RedirectUris,
		ResponseTypes:           request.ResponseTypes,
		GrantTypes:              request.GrantTypes,
		PostLogoutRedirectUris:  request.PostLogoutRedirectUris,
	}
	saveClient(client)

	return client
}

func mapToAddClientResponse(client Client) dto.AddClientResponse {
	return dto.AddClientResponse{
		ClientId:     client.ClientId,
		ClientSecret: client.ClientSecret,
	}
}

func generateRandomString() string {
	b := make([]rune, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
