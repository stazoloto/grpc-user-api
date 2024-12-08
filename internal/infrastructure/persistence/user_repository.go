package persistence

import (
	"errors"

	"github.com/stazoloto/auth_microservice/internal/app/domain"
)

type UserRepository struct {
	users map[string]*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	user.ID = generateID()
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) GetUserByID(id string) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *domain.User) error {
	_, exists := r.users[user.ID]
	if !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	_, exists := r.users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}
