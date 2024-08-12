package main

import (
	db "example.com/go-crud/config"
	"example.com/go-crud/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create the fiber app
	app := fiber.New()
	// Routers
	routes.Routes(app)
	// DB connection
	db.DatabaseConnection()
	// Server running on
	app.Listen(":3000")
}
