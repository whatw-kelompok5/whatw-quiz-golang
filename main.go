package main

import (
	"fmt"
	"log"
	"os"
	grpc "whatw-golang/cmd/server/gRPC"
	gateway "whatw-golang/cmd/server/gateway"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := fiber.New()
	PORTFIBER := os.Getenv("FIBER_SERVER_ADDRESS")

	go grpc.RunGRPCServer()
	go gateway.RunGRPCServerGateway()

	fmt.Println("Fiber server is running on port:"+ PORTFIBER)
	if err := app.Listen(":"+ PORTFIBER); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}