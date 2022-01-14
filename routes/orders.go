package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
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

func CreateOrder(c *fiber.Ctx) error { return nil }
func GetOrders(c *fiber.Ctx) error   { return nil }
func GetOrder(c *fiber.Ctx) error    { return nil }
func UpdateOrder(c *fiber.Ctx) error { return nil }
func DeleteOrder(c *fiber.Ctx) error { return nil }
