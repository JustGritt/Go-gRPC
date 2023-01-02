package routes

import (
	"fmt"
	"time"

	"github.com/JustGritt/go-grpc/database"
	"github.com/JustGritt/go-grpc/models"
	"github.com/gofiber/fiber/v2"
)

type Stream struct {
	Notifier               chan []byte
	newClients             chan chan []byte
	closeClientConnections chan chan []byte
	clients                map[chan []byte]bool
}

func GetStream(c *fiber.Ctx) error {
	var payments []models.Payment
	database.Database.Db.Find(&payments)

	if len(payments) == 0 {
		return c.Status(404).JSON("No payments found")
	}

	go func() {
		for {
			var newPayments []models.Payment
			database.Database.Db.Find(&newPayments)

			if len(newPayments) > len(payments) {
				payments = newPayments
				fmt.Println("New payment made: ", payments[len(payments)-1])

			}
			time.Sleep(1 * time.Second)
		}
	}()

	return c.Status(200).JSON(CreateResponsePayments(payments))
}
