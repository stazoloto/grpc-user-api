package grpc

import (
	"log"
	"net"

	"github.com/stazoloto/auth_microservice/internal/app/service"
	user "github.com/stazoloto/auth_microservice/pkg/proto/user"
	"google.golang.org/grpc"
)

func StartGRPCServer(userService *service.UserService, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, NewUserGRPCServer(userService))
	log.Printf("grpc server listening on %v", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}