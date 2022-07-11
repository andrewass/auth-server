package client

import (
	"auth-server/client/dto"
	"math/rand"
)

var clients = []Client{
	{ClientId: "firstId", ClientSecret: "firstSecret"},
	{ClientId: "secId", ClientSecret: "secSecret"},
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getClients() []Client {
	return getAllClients()
}

func addClient(request dto.AddClientRequest) dto.AddClientResponse {
	client := Client{
		ClientId:     generateRandomString(),
		ClientSecret: generateRandomString(),
	}
	saveClient(client)
	return mapToAddClientResponse(client)
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
