package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roryjarrard/fiber-api/database"
	"github.com/roryjarrard/fiber-api/models"
)

type User struct {
	// this is NOT the model User, use as serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(201).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	err := database.Database.Db.Find(&users).Error
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUsers := []User{}
	for _, u := range users {
		responseUsers = append(responseUsers, CreateResponseUser(u))
	}

	return c.Status(200).JSON(responseUsers)
}
