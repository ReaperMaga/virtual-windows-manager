package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Password  string             `bson:"password" json:"-"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

type LoginSession struct {
	Id       string             `bson:"_id"`
	UserId   primitive.ObjectID `bson:"user_id"`
	ExpireAt time.Time          `bson:"expire_at"`
}
