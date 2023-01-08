package authorization

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorizationRepository struct {
	Collection *mongo.Collection
}

func (r *AuthorizationRepository) saveAuthorizationCode(authCode AuthCode) {
	ctx := context.TODO()
	_, err := r.Collection.InsertOne(ctx, authCode)
	if err != nil {
		panic(err)
	}
}

func (r *AuthorizationRepository) deleteAuthorizationCode(code AuthCode) {
	ctx := context.TODO()
	_, err := r.Collection.DeleteOne(ctx, code)
	if err != nil {
		panic(err)
	}
}

func (r *AuthorizationRepository) getAuthorizationCode(code string) AuthCode {
	ctx := context.TODO()
	var authCode AuthCode
	err := r.Collection.FindOne(ctx, bson.M{"code": code}).Decode(&authCode)
	if err != nil {
		panic(err)
	}
	return authCode
}
