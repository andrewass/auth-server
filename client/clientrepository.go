package client

import (
	"auth-server/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const clientCollection = "client"

func getClientById(clientId string) *Client {
	ctx := context.TODO()
	response := common.Database.Collection(clientCollection).FindOne(ctx, bson.M{"client_id": clientId})

	return extractClientFromSingleResult(response)
}

func getAllClients(email string) []Client {
	ctx := context.TODO()
	var clients []Client
	response, err := common.Database.Collection(clientCollection).Find(ctx, bson.M{"user_email": email})
	if err != nil {
		panic(err)
	}
	if err = response.All(ctx, &clients); err != nil {
		panic(err)
	}
	return clients
}

func saveNewClient(client Client) {
	ctx := context.TODO()
	_, err := common.Database.Collection(clientCollection).InsertOne(ctx, client)
	if err != nil {
		panic(err)
	}
}

func saveUpdatedClient(client Client) {
	ctx := context.TODO()
	_, err := common.Database.Collection(clientCollection).ReplaceOne(ctx, bson.M{"_id": client.ID}, client)
	if err != nil {
		panic(err)
	}
}

func extractClientFromSingleResult(result *mongo.SingleResult) *Client {
	var storedClient Client
	if err := result.Decode(&storedClient); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	return &storedClient
}
