package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func main() {
	err = connectDB()
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
