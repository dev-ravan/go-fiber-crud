package main

import (
	db "example.com/go-crud/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create the fiber app
	app := fiber.New()
	// DB connection
	db.DatabaseConnection()
	// Server running on
	app.Listen(":3000")
}
