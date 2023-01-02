package client

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// AddAdminClient For testing during development only
func AddAdminClient() {
	adminClients := getClients("admin@admin.com")
	log.Println("Check if admin client exists")
	if len(adminClients) == 0 {
		log.Println("Attempting to create new admin client")
		createNewClient(AddClientRequest{
			UserEmail:    "admin@admin.com",
			ClientUri:    viper.Get("FRONTEND_URL").(string),
			ClientSecret: viper.Get("CLIENT_SECRET").(string),
			Type:         Internal,
		})
		log.Println("New admin client created")
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

func createNewClient(request AddClientRequest) Client {
	return persistNewClient(request)
}

func persistNewClient(request AddClientRequest) Client {
	hashedSecret, _ := bcrypt.GenerateFromPassword([]byte(request.ClientSecret), 8)
	client := Client{
		ID:                      primitive.NewObjectID(),
		ClientId:                generateRandomString(),
		ClientSecret:            string(hashedSecret),
		ClientIdIssuedAt:        time.Now(),
		ClientType:              request.Type,
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
	client.ClientSecret = request.ClientSecret

	return client
}

func updateClient(request UpdateClientRequest) Client {
	client := GetClientByIdAndSecret(request.ClientID, request.ClientSecret)
	saveUpdatedClient(client)
	return client
}

func deleteClient(clientId string) {
	deleteClientByClientId(clientId)
}

func rotateClientSecret(clientId string) Client {
	client := GetClientById(clientId)
	newSecret := generateRandomString()
	newHash, _ := bcrypt.GenerateFromPassword([]byte(newSecret), 8)
	client.ClientSecret = string(newHash)
	saveUpdatedClient(client)
	client.ClientSecret = newSecret

	return client
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
