package auth

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const salt int = 14

func CreateUser(name string, password string) (*User, error) {
	hashed, err := hash(password)
	if err != nil {
		return nil, err
	}
	user := &User{Id: primitive.NewObjectID(), Name: name, Password: hashed, CreatedAt: time.Now()}
	err = Repository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ValidatePassword(user *User, password string) bool {
	return compareHash(user.Password, password)
}

func hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), salt)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func compareHash(hashValue string, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashValue), []byte(value))
	return err == nil
}

func log(err error) {
	fmt.Println("There was an error in user: ", err)
}
