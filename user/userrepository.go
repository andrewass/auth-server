package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (r *UserRepository) saveUser(user User) {
	ctx := context.TODO()
	_, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		panic(err)
	}
}

func (r *UserRepository) findUserByEmail(email string) *User {
	ctx := context.TODO()
	response := r.Collection.FindOne(ctx, bson.M{"email": email})
	return r.extractUserFromSingleResult(response)
}

func (r *UserRepository) findUserBySubject(subject string) *User {
	ctx := context.TODO()
	response := r.Collection.FindOne(ctx, bson.M{"subject": subject})
	return r.extractUserFromSingleResult(response)
}

func (r *UserRepository) existsUserByEmail(email string) bool {
	ctx := context.TODO()
	response := r.Collection.FindOne(ctx, bson.M{"email": email})
	return r.extractUserFromSingleResult(response) != nil
}

func (r *UserRepository) extractUserFromSingleResult(result *mongo.SingleResult) *User {
	var storedUser User
	if err := result.Decode(&storedUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	return &storedUser
}
