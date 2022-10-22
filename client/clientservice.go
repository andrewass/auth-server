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

func VerifyClient(clientId string, clientSecret string) {

}

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

func generateRandomString() string {
	b := make([]rune, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
