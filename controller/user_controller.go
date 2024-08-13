package controller

import (
	"time"

	db "example.com/go-crud/config"
	models "example.com/go-crud/models"
	"github.com/gofiber/fiber/v2"
)

func ListOfUsers(c *fiber.Ctx) error {
	var user []models.User
	db.DB.Select("*").Find(&user)

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"data":   user,
	})
}

func SingleUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User

	db.DB.Select("*").Find(&user, userId)

	userDetails := make(map[string]interface{})
	userDetails["id"] = user.Id
	userDetails["name"] = user.Name
	userDetails["emailId"] = user.EmailId

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"data":   userDetails,
	})

}

// =========> Update user
func UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User

	db.DB.Find(&user, userId)

	// Validation
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"msg":    "User not found",
		})
	}

	var updatedUser models.User

	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
			"msg":    "Invalid data",
		})
	}

	if updatedUser.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": false,
			"msg":    "Name is a required field",
		})
	}

	user.Name = updatedUser.Name
	user.EmailId = updatedUser.EmailId
	db.DB.Save(&user)
	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"msg":    "User updated successfully",
	})
}

// =========> Create user
func CreateUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid data",
		})
	}

	if data["name"] == "" || data["emailId"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Add the required fields!",
		})
	}

	// Add user to the db
	user := models.User{
		Name:      data["name"],
		EmailId:   data["emailId"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{}}

	db.DB.Create(&user)
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "User added successfully..!",
		"data":    data,
	})

}

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	db.DB.Find(&user, id)

	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"msg":    "User not found",
		})
	}

	db.DB.Delete(&user)
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "User removed successfully..!",
	})

}
