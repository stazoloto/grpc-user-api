package service

import (
	"context"

	"github.com/stazoloto/auth_microservice/internal/app/domain"
	"github.com/stazoloto/auth_microservice/internal/app/usecase"
	user "github.com/stazoloto/auth_microservice/pkg/proto/user"
)

type UserService struct {
	usecase *usecase.UserUsecase
}

func NewUserService(usecase *usecase.UserUsecase) *UserService {
	return &UserService{
		usecase: usecase,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	newUser := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	err := s.usecase.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{Id: newUser.ID}, nil
}

func (s *UserService) GetUserByID(ctx context.Context, req *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error) {
	usr, err := s.usecase.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &user.GetUserByIDResponse{
		User: &user.User{
			Id:    usr.ID,
			Name:  usr.Name,
			Email: usr.Email,
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	updatedUser := &domain.User{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	}

	err := s.usecase.UpdateUser(updatedUser)
	if err != nil {
		return nil, err
	}

	return &user.UpdateUserResponse{}, nil

}

func (s *UserService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	err := s.usecase.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserResponse{}, nil
}
