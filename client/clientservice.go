package client

import (
	"auth-server/client/dto"
)

var clients = []Client{
	{ClientId: "firstId", ClientSecret: "firstSecret"},
	{ClientId: "secId", ClientSecret: "secSecret"},
}

func getClient(clientId string) Client {
	return findClient(clientId)
}

func getAllClients() []Client {
	return clients
}

func addClient(request dto.AddClientRequest) dto.AddClientResponse {
	saveClient(Client{ClientId: "2323fdsf"})
	return dto.AddClientResponse{}
}
