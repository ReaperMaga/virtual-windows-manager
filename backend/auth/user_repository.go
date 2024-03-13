package auth

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"virtual-windows-manager/database"
)

var Repository UserRepository

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

func (repo *MongoUserRepository) FindByNameOrErr(name string) (*User, error) {
	findResult := repo.Collection.FindOne(database.Context, bson.D{{"name", name}})
	err := findResult.Err()
	if err != nil {
		return nil, err
	}
	var result *User
	err = findResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *MongoUserRepository) ExistsByName(name string) bool {
	result := repo.Collection.FindOne(database.Context, bson.D{{"name", name}})
	return result.Err() == nil
}

func (repo *MongoUserRepository) Delete(user *User) bool {
	result, err := repo.Collection.DeleteOne(database.Context, bson.D{{"name", user.Name}})
	if err != nil {
		return false
	}
	return result.DeletedCount > 0
}

func (repo *MongoUserRepository) Update(user *User) (*User, error) {
	_, err := repo.Collection.ReplaceOne(database.Context, bson.D{{"name", user.Name}}, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *MongoUserRepository) Count() int64 {
	count, err := repo.Collection.CountDocuments(database.Context, bson.D{})
	if err != nil {
		return 0
	}
	return count
}

func (repo *MongoUserRepository) FindByIdOrErr(id primitive.ObjectID) (*User, error) {
	findResult := repo.Collection.FindOne(database.Context, bson.D{{"_id", id}})
	err := findResult.Err()
	if err != nil {
		return nil, err
	}
	var result *User
	err = findResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
