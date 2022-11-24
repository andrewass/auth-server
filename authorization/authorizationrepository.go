package authorization

import (
	"auth-server/common"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const authorizationCodeCollection = "authorizationCode"

func saveAuthorizationCode(authCode AuthCode) {
	ctx := context.TODO()
	_, err := common.Database.Collection(authorizationCodeCollection).InsertOne(ctx, authCode)
	if err != nil {
		panic(err)
	}
}

func deleteAuthorizationCode(code AuthCode) {
	ctx := context.TODO()
	_, err := common.Database.Collection(authorizationCodeCollection).DeleteOne(ctx, code)
	if err != nil {
		panic(err)
	}
}

func getAuthorizationCode(code string) AuthCode {
	ctx := context.TODO()
	var authCode AuthCode
	err := common.Database.Collection(authorizationCodeCollection).FindOne(ctx, bson.M{"code": code}).Decode(&authCode)
	if err != nil {
		panic(err)
	}
	return authCode
}
