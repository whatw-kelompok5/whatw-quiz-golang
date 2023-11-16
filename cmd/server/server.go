// server.go
package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"whatw-golang/cmd/server/laravel_grpc"
	pb_laravel "whatw-golang/pb/laravel"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func RunGRPCServer() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT_SERVER_GRPC")

	// gRPC server
	grpcServer := grpc.NewServer()
	grpc.ServiceRegistrar(grpcServer).RegisterService(&pb_laravel.QuestionService_ServiceDesc, pb_laravel.QuestionServiceServer(&laravel_grpc.Server{}))

	// Start gRPC server in a separate goroutine
	go func() {
		listener, err := net.Listen("tcp", ":"+PORT)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		fmt.Println("gRPC server is running on port" + PORT)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}
