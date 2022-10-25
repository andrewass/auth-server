package client

import (
	"auth-server/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const clientCollection = "client"

func getAllClients(email string) []Client {
	ctx := context.TODO()
	response, err := common.Database.Collection(clientCollection).Find(ctx, bson.M{"email": email})
	if err == nil {
		panic(err)
	}
	var clients []Client
	if err = response.Decode(clients); err != nil {
		panic(err)
	}
	return clients
}

func saveClient(client Client) {
	ctx := context.TODO()
	_, err := common.Database.Collection(clientCollection).InsertOne(ctx, client)
	if err != nil {
		panic(err)
	}
}
