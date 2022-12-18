package main

import (
	"github.com/JustGritt/go-grpc/database"
	"github.com/JustGritt/go-grpc/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber!")
}

func setupRoutes(app *fiber.App) {
	// GET routes
	// =================
	app.Get("/api", welcome) // Welcome
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Get("/api/payments", routes.GetPayments)
	app.Get("/api/payments/:id", routes.GetPayment)
	// app.Get("/api/payments/user/:id", routes.GetPaymentsByUser)
	// app.Get("/api/payments/product/:id", routes.GetPaymentsByProduct)

	// POST routes
	// =================
	app.Post("/api/users", routes.CreateUser)
	app.Post("/api/products", routes.CreateProduct)
	app.Post("/api/payments", routes.CreatePayment)

	// PUT routes
	// =================
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Put("/api/payments/:id", routes.UpdatePayment)

	// DELETE routes
	// =================
	app.Delete("/api/users/:id", routes.DeleteUser)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	app.Delete("/api/payments/:id", routes.DeletePayment)
}

func main() {
	database.Connect()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
