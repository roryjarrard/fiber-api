package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/roryjarrard/fiber-api/database"
	"github.com/roryjarrard/fiber-api/models"
)

// {id: 1, user: {id: 2, first_name: "Fred", last_name: "Jones"}, product: {id: 24, name: "Macbook", serial_number: "322348x"}}
type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRef, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRef, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(201).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error { return nil }

func GetOrder(c *fiber.Ctx) error { return nil }

func UpdateOrder(c *fiber.Ctx) error { return nil }

func DeleteOrder(c *fiber.Ctx) error { return nil }
