package main

import "github.com/gofiber/fiber/v2"

// type testApiJson struct {
// 	status  bool
// 	message string
// }

func main() {
	// Create the fiber app
	app := fiber.New()

	// Create route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello,Go World!")
	})

	testApiJson := fiber.Map{"status": true,
		"message": "Hello Go Fiber"}

	// Test get api
	app.Get("/testApi", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(testApiJson)
	})
	// Server running on
	app.Listen(":3000")
}
