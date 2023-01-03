package routes

import "github.com/gofiber/fiber/v2"

func SendMessages(c *fiber.Ctx) error {
	return c.SendString("Send messages")
}
