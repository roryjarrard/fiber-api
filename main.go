package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/roryjarrard/fiber-api/database"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected successfully")

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API")
	})

	log.Fatal(app.Listen(":1337"))
}
