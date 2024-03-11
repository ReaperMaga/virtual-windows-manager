package auth

type UserRepository interface {
	Create(user *User) error
	Delete(user *User) bool
	Update(user *User) (*User, error)
	FindByNameOrErr(name string) (*User, error)
	ExistsByName(name string) bool
}

var Repository UserRepository
