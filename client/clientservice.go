package client

import (
	"auth-server/client/dto"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// AddAdminClient For testing during development only
func AddAdminClient() {
	adminClients := getClients("admin@admin.com")
	if len(adminClients) == 0 {
		adminClient := addClient(dto.AddClientRequest{
			UserEmail: "admin@admin.com",
			ClientUri: viper.Get("FRONTEND_URL").(string),
		})
		adminClient.ClientId = viper.Get("CLIENT_ID").(string)
		adminClient.ClientSecret = viper.Get("CLIENT_SECRET").(string)
		adminClient.ClientType = Internal
		saveUpdatedClient(adminClient)
	}
}

func GetClientById(clientId string) Client {
	return *getClientById(clientId)
}

func GetClientByIdAndSecret(clientId string, clientSecret string) Client {
	client := getClientById(clientId)
	verifyMatchingSecret(clientSecret, client.ClientSecret)
	return *client
}

func getClients(email string) []Client {
	return getAllClients(email)
}

func addClient(request dto.AddClientRequest) Client {
	client := Client{
		ID:                      primitive.NewObjectID(),
		ClientId:                generateRandomString(),
		ClientSecret:            generateRandomString(),
		ClientIdIssuedAt:        time.Now(),
		ClientType:              External,
		UserEmail:               request.UserEmail,
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
	saveNewClient(client)

	return client
}

func updateClient(request dto.UpdateClientRequest) Client {
	client := GetClientByIdAndSecret(request.ClientID, request.ClientSecret)
	saveUpdatedClient(client)
	return client
}

func deleteClient(clientId string, clientSecret string) {
	client := *getClientById(clientId)
	verifyMatchingSecret(clientSecret, client.ClientSecret)
}

func verifyMatchingSecret(plaintextSecret, hashedSecret string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedSecret), []byte(plaintextSecret))
	if err != nil {
		panic("Mismatch of client secret")
	}
}

func generateRandomString() string {
	b := make([]rune, 30)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
