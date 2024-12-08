package domain

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}
