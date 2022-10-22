package authorization

import (
	"auth-server/common"
	"context"
)

const authorizationCodeCollection = "authorizationCode"

func saveAuthorizationCode(authCode AuthCode) {
	ctx := context.TODO()
	_, err := common.Database.Collection(authorizationCodeCollection).InsertOne(ctx, authCode)
	if err != nil {
		panic(err)
	}
}
