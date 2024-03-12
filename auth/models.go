package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
}

type LoginSession struct {
	Id       string             `bson:"_id"`
	UserId   primitive.ObjectID `bson:"user_id"`
	ExpireAt time.Time          `bson:"expire_at"`
}
