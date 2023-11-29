package cmd

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"whatw-golang/pb/laravel/avatar"
	"whatw-golang/pb/laravel/question"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"golang.org/x/net/http2"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8081", "gRPC server endpoint")
)

func RunGRPCServerGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := question.RegisterQuestionServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	if err := avatar.RegisterAvatarServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	// Configure HTTP/2 settings
	server := &http.Server{
		Addr:    ":8083", // Set the desired port
		Handler: mux,
	}

	http2.ConfigureServer(server, &http2.Server{
		MaxConcurrentStreams: 100,
	})

	fmt.Printf("gRPC Gateway server is running on :8083 (connecting to gRPC server at %s)\n", *grpcServerEndpoint)
	return server.ListenAndServe()
}
