package controller

import (
	db "example.com/go-crud/config"
	models "example.com/go-crud/models"
	"github.com/gofiber/fiber/v2"
)

func ListOfUsers(c *fiber.Ctx) error {
	var usersList []models.User
	db.DB.Find(&usersList)
	return c.JSON(usersList)
}

func SingleUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	db.DB.Find(&user, id)
	return c.JSON(&user)

}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user = new(models.User)
	db.DB.First(&user, id)
	if user.EmailId == "" {
		return c.Status(500).SendString("User not available")
	}
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.DB.Save(&user)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.DB.Create(&user)
	return c.JSON(&user)

}

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	db.DB.First(&user, id)
	if user.EmailId == "" {
		return c.Status(500).SendString("User not available")
	}
	db.DB.Delete(&user)
	return c.SendString("User successfully deleted..!")

}
