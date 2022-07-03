package client

var clients = []Client{
	{ClientId: "firstId", ClientSecret: "firstSecret"},
	{ClientId: "secId", ClientSecret: "secSecret"},
}

func GetClient(clientId string) Client {
	return clients[0]
}

func GetAllClients() []Client {
	return clients
}

func AddClient() {

}
