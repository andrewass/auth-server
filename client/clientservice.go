package client

import (
	"auth-server/client/dto"
	"github.com/spf13/viper"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// AddAdminClient For testing during development only
func AddAdminClient() {
	adminClients := getClients("admin@admin.com")
	if len(adminClients) == 0 {
		adminClient := addClient(dto.AddClientRequest{
			UserEmail:  "admin@admin.com",
			ClientName: "adminClient",
			ClientUri:  "http://localhost:8090",
		})
		adminClient.ClientId = viper.Get("CLIENT_ID").(string)
		adminClient.ClientSecret = viper.Get("CLIENT_ID").(string)
		adminClient.ClientType = internal
	}
}

func VerifyClient(clientId string, clientSecret string) {
	client := getClientById(clientId)
	if client == nil {
		panic("Client not found")
	}
	if client.ClientSecret != clientSecret {
		panic("Received client secret is incorrect")
	}
}

func getClients(email string) []Client {
	return getAllClients(email)
}

func addClient(request dto.AddClientRequest) Client {
	client := Client{
		ClientId:                generateRandomString(),
		ClientSecret:            generateRandomString(),
		ClientIdIssuedAt:        time.Now(),
		ClientType:              external,
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

func updateClient(request dto.UpdateClientRequest) Client {
	client := getClientById(request.ClientID)
	return *client
}

func generateRandomString() string {
	b := make([]rune, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
