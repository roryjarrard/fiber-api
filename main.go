package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/roryjarrard/fiber-api/database"
	"github.com/roryjarrard/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected successfully")

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":1337"))
}
