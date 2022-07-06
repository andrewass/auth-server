package common

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Database *mongo.Database

func ConnectToDatabase() {
	credentials := options.Credential{
		Username: viper.Get("MONGO_DB_USERNAME").(string),
		Password: viper.Get("MONGO_DB_PASSWORD").(string),
	}
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credentials)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	ping(client, ctx)
	createDatabase(client)
}

func ping(client *mongo.Client, ctx context.Context) {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("MongoDB - connected successfully")
}

func createDatabase(client *mongo.Client) {
	Database = client.Database("authDatabase")
	Database.Collection("client")
}
