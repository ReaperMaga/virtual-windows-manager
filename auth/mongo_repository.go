package auth

import (
	"go.mongodb.org/mongo-driver/mongo"
	"virtual-windows-manager/database"
)

func NewMongoUserRepository() *MongoUserRepository {
	collection := database.Database.Collection("users")
	return &MongoUserRepository{
		Collection: collection,
	}
}

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func (repo *MongoUserRepository) Create(user *User) error {
	_, err := repo.Collection.InsertOne(database.Context, user)
	return err
}
