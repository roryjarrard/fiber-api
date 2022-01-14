package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roryjarrard/fiber-api/models"
)

// Product is a serializer, not a DB model
type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func GetProducts(c *fiber.Ctx) error   { return nil }
func GetProduct(c *fiber.Ctx) error    { return nil }
func CreateProduct(c *fiber.Ctx) error { return nil }
func UpdateProduct(c *fiber.Ctx) error { return nil }
func DeleteProduct(c *fiber.Ctx) error { return nil }
