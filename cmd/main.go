package main

import (
	"github.com/stazoloto/auth_microservice/internal/app/service"
	"github.com/stazoloto/auth_microservice/internal/app/usecase"
	"github.com/stazoloto/auth_microservice/internal/infrastructure/grpc"
	"github.com/stazoloto/auth_microservice/internal/infrastructure/persistence"
)

func main() {
	repo := persistence.NewUserRepository()
	usecase := usecase.NewUserUsecase(repo)
	userService := service.NewUserService(usecase)

	port := ":50051"
	grpc.StartGRPCServer(userService, port)
}
