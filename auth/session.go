package auth

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"virtual-windows-manager/database"
)

var SessionRepository LoginSessionRepository

type MongoLoginSessionRepository struct {
	Collection *mongo.Collection
}

func NewMongoLoginSessionRepository() *MongoLoginSessionRepository {
	collection := database.Database.Collection("sessions")
	return &MongoLoginSessionRepository{
		Collection: collection,
	}
}

func (repo *MongoLoginSessionRepository) Create(session *LoginSession) error {
	_, err := repo.Collection.InsertOne(database.Context, session)
	return err
}

func (repo *MongoLoginSessionRepository) FindByIdOrErr(id string) (*LoginSession, error) {
	findResult := repo.Collection.FindOne(database.Context, bson.D{{"_id", id}})
	err := findResult.Err()
	if err != nil {
		return nil, err
	}
	var result *LoginSession
	err = findResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *MongoLoginSessionRepository) ExistsById(id string) bool {
	result := repo.Collection.FindOne(database.Context, bson.D{{"_id", id}})
	return result.Err() == nil
}

func (repo *MongoLoginSessionRepository) Delete(session *LoginSession) bool {
	result, err := repo.Collection.DeleteOne(database.Context, bson.D{{"_id", session.Id}})
	if err != nil {
		return false
	}
	return result.DeletedCount > 0
}

func (repo *MongoLoginSessionRepository) Update(session *LoginSession) (*LoginSession, error) {
	_, err := repo.Collection.ReplaceOne(database.Context, bson.D{{"_id", session.Id}}, session)
	if err != nil {
		return nil, err
	}
	return session, nil
}
