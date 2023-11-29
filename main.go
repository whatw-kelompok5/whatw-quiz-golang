package main

import (
	"fmt"
	"log"
	"os"
	"whatw-golang/cmd"
	"whatw-golang/cmd/server"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := fiber.New()
	PORTFIBER := os.Getenv("FIBER_SERVER_ADDRESS")

	go server.RunGRPCServer()
	go cmd.RunGRPCServerGateway()

	fmt.Println("Fiber server is running on port:"+ PORTFIBER)
	if err := app.Listen(":"+ PORTFIBER); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}