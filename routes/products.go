package routes

import (
	"github.com/JustGritt/go-grpc/database"
	"github.com/JustGritt/go-grpc/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func CreateResponseProduct(product models.Product) Product {
	return Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if product.Name == "" {
		return c.Status(400).JSON("Product name is required")
	}

	var existingProduct models.Product
	database.Database.Db.Where("name = ?", product.Name).First(&existingProduct)
	if existingProduct.ID != 0 {
		return c.Status(400).JSON("Product name already taken")
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProductId(id int, product *models.Product) error {
	database.Database.Db.Find(&product, id)
	if product.ID == 0 {
		return fiber.NewError(404, "Product not found")
	}

	return nil
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.Database.Db.Find(&products)

	var responseProducts []Product
	for _, product := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(product))
	}

	return c.Status(200).JSON(responseProducts)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := GetProductId(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := GetProductId(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON("Invalid request")
	}

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := GetProductId(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Delete(&product)

	return c.Status(200).JSON("Product deleted")
}
