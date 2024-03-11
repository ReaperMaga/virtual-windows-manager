package auth

type UserRepository interface {
	Create(user *User) error
}

var Repository UserRepository
