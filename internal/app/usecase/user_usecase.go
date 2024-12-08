package usecase

import "github.com/stazoloto/auth_microservice/internal/app/domain"

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserUsecase) UpdateUser(user *domain.User) error {
	return u.repo.UpdateUser(user)
}

func (u *UserUsecase) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}
