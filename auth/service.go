package auth

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const salt int = 14

func Initialize() {
	Repository = NewMongoUserRepository()
	SessionRepository = NewMongoLoginSessionRepository()
}

func IsAuth(sessionToken string) bool {
	if sessionToken == "" {
		return false
	}
	session, err := SessionRepository.FindByIdOrErr(sessionToken)
	if err != nil {
		return false
	}
	if time.Now().After(session.ExpireAt) {
		SessionRepository.Delete(session)
		return false
	}
	return true
}

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

func CreateSession(user *User) (*LoginSession, error) {
	session := &LoginSession{
		Id:       uuid.NewString(),
		UserId:   user.Id,
		ExpireAt: time.Now().Add(time.Second * 20),
	}
	err := SessionRepository.Create(session)
	if err != nil {
		return nil, err
	}
	return session, nil
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
