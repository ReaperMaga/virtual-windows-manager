package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRepository interface {
	Create(user *User) error
	Delete(user *User) bool
	Update(user *User) (*User, error)
	FindByIdOrErr(id primitive.ObjectID) (*User, error)
	FindByNameOrErr(name string) (*User, error)
	ExistsByName(name string) bool
	Count() int64
}

type LoginSessionRepository interface {
	Create(session *LoginSession) error
	Delete(session *LoginSession) bool
	Update(session *LoginSession) (*LoginSession, error)
	FindByIdOrErr(id string) (*LoginSession, error)
	ExistsById(id string) bool
}
