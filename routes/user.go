package routes

import (
	"errors"

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

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("userId")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :userId is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("userId")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :userId is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if updateData.FirstName != "" {
		user.FirstName = updateData.FirstName
	}

	if updateData.LastName != "" {
		user.LastName = updateData.LastName
	}

	database.Database.Db.Save(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("userId")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :userId is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user, "id = ?", user.ID).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("successfully deleted user")
}
