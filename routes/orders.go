package routes

import (
	"errors"
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

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user, "id = ?", order.UserRef)
		database.Database.Db.Find(&product, "id = ?", order.ProductRef)
		responseOrders = append(responseOrders, CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product)))
	}

	return c.Status(200).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	order := models.Order{}
	user := models.User{}
	product := models.Product{}

	id, err := c.ParamsInt("orderId")
	if err != nil {
		return c.Status(400).JSON("provide an integer value")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	findUser(order.UserRef, &user)
	findProduct(order.ProductRef, &product)

	responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
	return c.Status(200).JSON(responseOrder)
}

func UpdateOrder(c *fiber.Ctx) error { return nil }

func DeleteOrder(c *fiber.Ctx) error { return nil }
