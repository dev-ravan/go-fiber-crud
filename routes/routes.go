package routes

import (
	"example.com/go-crud/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/users", controller.ListOfUsers)
	app.Get("/user/:id", controller.SingleUser)
	app.Put("/user/:id", controller.UpdateUser)
	app.Post("/user", controller.CreateUser)
	app.Delete("/user/:id", controller.DeleteUser)
}
