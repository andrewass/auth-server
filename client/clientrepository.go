package client

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	Collection *mongo.Collection
}

func (r *ClientRepository) getClientById(clientId string) *Client {
	ctx := context.TODO()
	response := r.Collection.FindOne(ctx, bson.M{"client_id": clientId})

	return r.extractClientFromSingleResult(response)
}

func (r *ClientRepository) getAllClients(email string) []Client {
	ctx := context.TODO()
	var clients []Client
	response, err := r.Collection.Find(ctx, bson.M{"user_email": email})
	if err != nil {
		panic(err)
	}
	if err = response.All(ctx, &clients); err != nil {
		panic(err)
	}
	return clients
}

func (r *ClientRepository) saveNewClient(client Client) {
	ctx := context.TODO()
	_, err := r.Collection.InsertOne(ctx, client)
	if err != nil {
		panic(err)
	}
}

func (r *ClientRepository) saveUpdatedClient(client Client) {
	ctx := context.TODO()
	_, err := r.Collection.ReplaceOne(ctx, bson.M{"_id": client.ID}, client)
	if err != nil {
		panic(err)
	}
}

func (r *ClientRepository) deleteClientByClientId(clientId string) {
	ctx := context.TODO()
	_, err := r.Collection.DeleteOne(ctx, bson.M{"client_id": clientId})
	if err != nil {
		panic("Unable to delete client document : " + err.Error())
	}
}

func (r *ClientRepository) extractClientFromSingleResult(result *mongo.SingleResult) *Client {
	var storedClient Client
	if err := result.Decode(&storedClient); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	return &storedClient
}
