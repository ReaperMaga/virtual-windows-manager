package vw

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"virtual-windows-manager/database"
)

var Repository VirtualWindowsRepository

type VirtualWindowsRepository interface {
	Create(vw *VirtualWindows) error
	Delete(vw *VirtualWindows) bool
	Update(vw *VirtualWindows) (*VirtualWindows, error)
	GetAll() ([]*VirtualWindows, error)
	FindByNameOrErr(name string) (*VirtualWindows, error)
	FindByIdOrErr(id string) (*VirtualWindows, error)
	ExistsByName(name string) bool
	Count() int64
}

func NewMongoVirtualWindowsRepository() *MongoVirtualWindowsRepository {
	collection := database.Database.Collection("vws")
	return &MongoVirtualWindowsRepository{
		Collection: collection,
	}
}

type MongoVirtualWindowsRepository struct {
	Collection *mongo.Collection
}

func (repo *MongoVirtualWindowsRepository) Create(vw *VirtualWindows) error {
	_, err := repo.Collection.InsertOne(database.Context, vw)
	return err
}

func (repo *MongoVirtualWindowsRepository) FindByNameOrErr(name string) (*VirtualWindows, error) {
	findResult := repo.Collection.FindOne(database.Context, bson.D{{"name", name}})
	err := findResult.Err()
	if err != nil {
		return nil, err
	}
	var result *VirtualWindows
	err = findResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *MongoVirtualWindowsRepository) FindByIdOrErr(id string) (*VirtualWindows, error) {
	findResult := repo.Collection.FindOne(database.Context, bson.D{{"_id", id}})
	err := findResult.Err()
	if err != nil {
		return nil, err
	}
	var result *VirtualWindows
	err = findResult.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *MongoVirtualWindowsRepository) ExistsByName(name string) bool {
	result := repo.Collection.FindOne(database.Context, bson.D{{"name", name}})
	return result.Err() == nil
}

func (repo *MongoVirtualWindowsRepository) Delete(vw *VirtualWindows) bool {
	result, err := repo.Collection.DeleteOne(database.Context, bson.D{{"_id", vw.Id}})
	if err != nil {
		return false
	}
	return result.DeletedCount > 0
}

func (repo *MongoVirtualWindowsRepository) Update(vw *VirtualWindows) (*VirtualWindows, error) {
	_, err := repo.Collection.ReplaceOne(database.Context, bson.D{{"name", vw.Name}}, vw)
	if err != nil {
		return nil, err
	}
	return vw, nil
}

func (repo *MongoVirtualWindowsRepository) Count() int64 {
	count, err := repo.Collection.CountDocuments(database.Context, bson.D{})
	if err != nil {
		return 0
	}
	return count
}

func (repo *MongoVirtualWindowsRepository) GetAll() ([]*VirtualWindows, error) {
	cursor, err := repo.Collection.Find(database.Context, bson.D{})
	if err != nil {
		return nil, err
	}
	var results []*VirtualWindows
	err = cursor.All(database.Context, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
