package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

var Context = context.TODO()

func Connect(uri string, database string) error {
	client, err := mongo.Connect(Context, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	err = client.Ping(Context, nil)
	if err != nil {
		return err
	}
	Database = client.Database(database)
	return nil
}
