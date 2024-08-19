package controller

import (
	"fmt"
	"time"

	db "example.com/go-crud/config"
	models "example.com/go-crud/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
}

func UserResponse(userModel models.User) User {
	return User{Id: userModel.Id, Name: userModel.Name, EmailId: userModel.EmailId}
}

// List of controllers
var ListOfUsers = listOfUsers
var CreateUser = createUser
var DeleteUser = removeUser
var UpdateUser = updateUser
var SingleUser = singleUser

func listOfUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)

	var responseUsers []User

	for _, user := range users {
		responseUser := UserResponse(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"data":   responseUsers,
	})
}

func singleUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	var user models.User

	db.DB.Select("*").Find(&user, userId)
	// Validation
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": false,
			"msg":    "User not found",
		})
	}

	userDetail := UserResponse(user)

	return c.Status(200).JSON(fiber.Map{
		"status": true,
		"data":   userDetail,
	})

}

// =========> Update user
func updateUser(c *fiber.Ctx) error {
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
func createUser(c *fiber.Ctx) error {
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

	// Check if email already exists
	var emailExists bool
	result := db.DB.Model(&models.User{}).Select("count(*) > 0").Where("email_id = ?", data["emailId"]).Find(&emailExists)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Error checking email",
		})
	}
	fmt.Println(emailExists, result)
	if emailExists {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Email already exists",
		})
	}

	// Add user to the db
	user := models.User{
		Name:      data["name"],
		EmailId:   data["emailId"],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to create user",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "User added successfully!",
		"data":    user,
	})
}

func removeUser(c *fiber.Ctx) error {
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
