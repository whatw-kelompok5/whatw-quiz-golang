// server.go
package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"whatw-golang/cmd/server/laravel_grpc/avatar"
	"whatw-golang/cmd/server/laravel_grpc/question"
	pb_laravel_avatar "whatw-golang/pb/laravel/avatar"
	pb_laravel_question "whatw-golang/pb/laravel/question"

	"google.golang.org/grpc"
)

func RunGRPCServer() {
	PORT := os.Getenv("PORT_SERVER_GRPC")

	// gRPC server
	grpcServer := grpc.NewServer()
	grpc.ServiceRegistrar(grpcServer).RegisterService(&pb_laravel_question.QuestionService_ServiceDesc, pb_laravel_question.QuestionServiceServer(&question.ServerQuestion{}))
	grpc.ServiceRegistrar(grpcServer).RegisterService(&pb_laravel_avatar.AvatarService_ServiceDesc, pb_laravel_avatar.AvatarServiceServer(&avatar.ServerAvatar{}))

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
