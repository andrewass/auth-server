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

type ClientService struct {
	Repository *ClientRepository
}

// AddAdminClient For testing during development only
func (s *ClientService) AddAdminClient() {
	adminClients := s.getClients("admin@admin.com")
	log.Println("Check if admin client exists")
	if len(adminClients) == 0 {
		log.Println("Attempting to create new admin client")
		s.createNewClient(AddClientRequest{
			UserEmail:    "admin@admin.com",
			ClientUri:    viper.Get("FRONTEND_URL").(string),
			ClientSecret: viper.Get("CLIENT_SECRET").(string),
			Type:         Internal,
		})
		log.Println("New admin client created")
	}
}

func (s *ClientService) GetClientById(clientId string) Client {
	return *s.Repository.getClientById(clientId)
}

func (s *ClientService) GetClientByIdAndSecret(clientId string, clientSecret string) Client {
	client := s.Repository.getClientById(clientId)
	s.verifyMatchingSecret(clientSecret, client.ClientSecret)
	return *client
}

func (s *ClientService) getClients(email string) []Client {
	return s.Repository.getAllClients(email)
}

func (s *ClientService) createNewClient(request AddClientRequest) Client {
	return s.persistNewClient(request)
}

func (s *ClientService) persistNewClient(request AddClientRequest) Client {
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
	s.Repository.saveNewClient(client)
	client.ClientSecret = request.ClientSecret

	return client
}

func (s *ClientService) updateClient(request UpdateClientRequest) Client {
	client := s.GetClientByIdAndSecret(request.ClientID, request.ClientSecret)
	s.Repository.saveUpdatedClient(client)
	return client
}

func (s *ClientService) deleteClient(clientId string) {
	s.Repository.deleteClientByClientId(clientId)
}

func (s *ClientService) rotateClientSecret(clientId string) Client {
	client := s.GetClientById(clientId)
	newSecret := generateRandomString()
	newHash, _ := bcrypt.GenerateFromPassword([]byte(newSecret), 8)
	client.ClientSecret = string(newHash)
	s.Repository.saveUpdatedClient(client)
	client.ClientSecret = newSecret

	return client
}

func (s *ClientService) verifyMatchingSecret(plaintextSecret, hashedSecret string) {
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
