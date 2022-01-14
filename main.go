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
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:userId", routes.GetUser)
	app.Put("/api/users/:userId", routes.UpdateUser)
	app.Delete("/api/users/:userId", routes.DeleteUser)

	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:productId", routes.GetProduct)
	app.Put("/api/products/:productId", routes.UpdateProduct)
	app.Delete("/api/products/:productId", routes.DeleteProduct)

	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:orderId", routes.GetOrder)
	app.Delete("/api/orders/:orderId", routes.DeleteOrder)
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
