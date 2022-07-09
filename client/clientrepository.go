package client

import (
	"auth-server/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const clientCollection = "client"

func findClient(clientId string) Client {
	ctx := context.TODO()
	var client Client
	err := common.Database.Collection(clientCollection).FindOne(ctx, bson.M{"client_id": clientId}).Decode(client)
	if err == nil {
		panic(err)
	}
	return client
}

func saveClient(client Client) {
	ctx := context.TODO()
	_, err := common.Database.Collection(clientCollection).InsertOne(ctx, client)
	if err != nil {
		panic(err)
	}
}
