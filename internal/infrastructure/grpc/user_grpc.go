package grpc

import (
	"context"

	"github.com/stazoloto/auth_microservice/internal/app/service"
	user "github.com/stazoloto/auth_microservice/pkg/proto/user"
) 

type UserGRPCServer struct {
	user.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewUserGRPCServer(userService *service.UserService) *UserGRPCServer {
	return &UserGRPCServer{
		userService: userService,
	}
} 

func (s *UserGRPCServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return s.userService.CreateUser(ctx, req)
} 

func (s *UserGRPCServer) GetUserByID(ctx context.Context, req *user.GetUserByIDRequest) (*user.GetUserByIDResponse, error) {
	return s.userService.GetUserByID(ctx, req)
} 

func (s *UserGRPCServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	return s.userService.UpdateUser(ctx, req)
} 

func (s *UserGRPCServer) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	return s.userService.DeleteUser(ctx, req)
} 

