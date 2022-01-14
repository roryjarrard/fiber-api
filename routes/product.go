package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roryjarrard/fiber-api/database"
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

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	if err := database.Database.Db.Find(&products).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var responseProducts []Product = []Product{}

	for _, p := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(p))
	}

	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("productId")
	if err != nil {
		return c.Status(400).JSON("please supply an integer value")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("productId")
	if err != nil {
		return c.Status(400).JSON("please supply an integer value")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if updateData.Name != "" {
		product.Name = updateData.Name
	}

	if updateData.SerialNumber != "" {
		product.SerialNumber = updateData.SerialNumber
	}

	database.Database.Db.Save(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(201).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error { return nil }
