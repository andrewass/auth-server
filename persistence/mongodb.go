package persistence

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDatabase() {
	credentials := options.Credential{
		Username: "mongoUser",
		Password: "mongoPassword",
	}
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	ping(client, ctx)
	createDatabase(ctx, client)
}

func ping(client *mongo.Client, ctx context.Context) {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB - connected successfully")
}

func createDatabase(ctx context.Context, client *mongo.Client) {
	database := client.Database("authDatabase")
	database.Collection("clients")
}
