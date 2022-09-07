package user

import (
	"auth-server/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "user"

func saveUser(user User) string {
	ctx := context.TODO()
	result, err := common.Database.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(string)
}

func findUserByEmail(email string) *User {
	ctx := context.TODO()
	response := common.Database.Collection(userCollection).FindOne(ctx, bson.M{"email": email})
	return extractUserFromSingleResult(response)
}

func existsUserByEmail(email string) bool {
	ctx := context.TODO()
	response := common.Database.Collection(userCollection).FindOne(ctx, bson.M{"email": email})
	return extractUserFromSingleResult(response) != nil
}

func extractUserFromSingleResult(result *mongo.SingleResult) *User {
	var storedUser User
	if err := result.Decode(&storedUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	return &storedUser
}
