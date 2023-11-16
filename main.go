package main

import (
	"fmt"
	"log"
	"os"
	"whatw-golang/cmd/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	PORTFIBER := os.Getenv("PORT_SERVER_FIBER")

	server.RunGRPCServer()

	fmt.Println("Fiber server is running on port" + "", PORTFIBER)
	if err := app.Listen(PORTFIBER); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}