package auth

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const salt int = 14

func Initialize() {
	Repository = NewMongoUserRepository()
	SessionRepository = NewMongoLoginSessionRepository()

	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if Repository.Count() == 0 {
		_, err := CreateUser(adminUsername, adminPassword)
		if err != nil {
			panic("There was an error while trying to create the default user: " + err.Error())
		}
	}
}

func IsAuthAndGetUser(username string, password string) (bool, *User) {
	user, err := Repository.FindByNameOrErr(username)
	if err != nil {
		return false, nil
	}
	if compareHash(user.Password, password) {
		return true, user
	}
	return false, nil
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

func GetSession(sessionToken string) (*LoginSession, error) {
	if sessionToken == "" {
		return nil, errors.New("session token is empty")
	}
	session, err := SessionRepository.FindByIdOrErr(sessionToken)
	if err != nil {
		return nil, err
	}
	if time.Now().After(session.ExpireAt) {
		SessionRepository.Delete(session)
		return nil, errors.New("session timeout")
	}
	return session, nil
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
	sessionTime, err := strconv.ParseInt(os.Getenv("SESSION_TIME"), 10, 64)
	if err != nil {
		return nil, err
	}
	session := &LoginSession{
		Id:       uuid.NewString(),
		UserId:   user.Id,
		ExpireAt: time.Now().Add(time.Second * time.Duration(sessionTime)),
	}
	err = SessionRepository.Create(session)
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
