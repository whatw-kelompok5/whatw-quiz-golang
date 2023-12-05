// server.go
package server

import (
	"fmt"
	"log"
	"net"
	"os"
	avatar "whatw-golang/cmd/handlers/avatar"
	question "whatw-golang/cmd/handlers/question"
	pb_laravel_avatar "whatw-golang/pb/laravel/avatar"
	pb_laravel_question "whatw-golang/pb/laravel/question"

	"google.golang.org/grpc"
)

func RunGRPCServer() {
	PORT := os.Getenv("GRPC_SERVER_ADDRESS")
	
	// gRPC server
	grpcServer := grpc.NewServer()
	pb_laravel_avatar.RegisterAvatarServiceServer(grpcServer, &avatar.ServerAvatar{})
	pb_laravel_question.RegisterQuestionServiceServer(grpcServer, &question.ServerQuestion{})

	// Start gRPC server in a separate goroutine
	go func() {
		listener, err := net.Listen("tcp", ":"+PORT)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		fmt.Println("gRPC server is running on port:" + PORT)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}