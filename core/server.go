package core

import (
	"log"
	"net"

	userPb "github.com/prayogatriady/sawer-grpc/model"
	"github.com/prayogatriady/sawer-grpc/services"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const (
	GRPC_PORT = ":50051"
)

func InitializeGRPCServer(db *gorm.DB) {
	// userRepo := repository.NewUserRepository(db)
	// userServ := services.NewUserService(userRepo)

	netListen, err := net.Listen("tcp", GRPC_PORT)
	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	userService := &services.UserService{DB: db}
	userPb.RegisterUserServiceServer(grpcServer, userService)

	log.Printf("Server started at %v", netListen.Addr())

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
