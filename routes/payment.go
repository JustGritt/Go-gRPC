package routes

import (
	"github.com/JustGritt/go-grpc/database"
	"github.com/JustGritt/go-grpc/models"
	"github.com/gofiber/fiber/v2"
)

type Payment struct {
	ID        uint `json:"id"`
	ProductId uint `json:"product_id"`
	Amount    uint `json:"amount"`
}

func CreateResponsePayment(payment models.Payment) Payment {
	return Payment{
		ID:        payment.ID,
		ProductId: payment.ProductID,
		Amount:    payment.Price,
	}
}

func CreateResponsePayments(payments []models.Payment) []Payment {
	var response []Payment
	for _, payment := range payments {
		response = append(response, CreateResponsePayment(payment))
	}
	return response
}

func GetPayments(c *fiber.Ctx) error {
	var payments []models.Payment
	database.Database.Db.Find(&payments)

	if len(payments) == 0 {
		return c.Status(404).JSON("No payments found")
	}

	return c.Status(200).JSON(CreateResponsePayments(payments))
}

func GetPayment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var payment models.Payment
	database.Database.Db.Where("id = ?", id).Find(&payment)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if payment.ID == 0 {
		return c.Status(404).JSON("Payment not found")
	}

	return c.Status(200).JSON(CreateResponsePayment(payment))
}

func CreatePayment(c *fiber.Ctx) error {
	var payment models.Payment

	if err := c.BodyParser(&payment); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&payment)
	return c.Status(200).JSON(CreateResponsePayment(payment))
}

func DeletePayment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var payment models.Payment
	database.Database.Db.Where("id = ?", id).Find(&payment)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if payment.ID == 0 {
		return c.Status(404).JSON("Payment not found")
	}

	database.Database.Db.Delete(&payment)
	return c.Status(200).JSON(CreateResponsePayment(payment))
}

// Update the payment with the given id
func UpdatePayment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var payment models.Payment
	database.Database.Db.Where("id = ?", id).Find(&payment)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if payment.ID == 0 {
		return c.Status(404).JSON("Payment not found")
	}

	if err := c.BodyParser(&payment); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Save(&payment)
	return c.Status(200).JSON(CreateResponsePayment(payment))
}

func GetPaymentByProductId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var payments []models.Payment
	database.Database.Db.Where("product_id = ?", id).Find(&payments)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if len(payments) == 0 {
		return c.Status(404).JSON("Payment not found")
	}

	return c.Status(200).JSON(CreateResponsePayments(payments))
}

func GetAllPaymentsByProductId(c *fiber.Ctx) error {
	var payments []models.Payment
	database.Database.Db.Find(&payments)

	if len(payments) == 0 {
		return c.Status(404).JSON("No payments found")
	}

	return c.Status(200).JSON(CreateResponsePayments(payments))
}
