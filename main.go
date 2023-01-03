package main

import (
	"github.com/JustGritt/go-grpc/database"
	"github.com/JustGritt/go-grpc/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	jwtware "github.com/gofiber/jwt/v3"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber!")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	c.SendString("Welcome " + name + "!")
	return c.Next()
}

func setupRoutes(app *fiber.App) {
	// Welcome
	app.Get("/api", welcome)

	// Login route
	app.Post("/login", routes.Login)

	// Stream routes
	// =================
	app.Get("/api/stream", routes.GetStream)

	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Token is missing, returns with error code 400 "Missing or malformed JWT"
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
		},
		SigningKey: []byte("secret"),
	}))

	// Restricted Routes
	app.Get("/restricted", restricted)

	// GET routes
	// =================
	app.Get("/api", welcome) // Welcome
	app.Get("/api/routess", routes.GetUsers)
	app.Get("/api/routess/:id", routes.GetUser)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Get("/api/payments", routes.GetPayments)
	app.Get("/api/payments/:id", routes.GetPayment)
	// app.Get("/api/payments/routes/:id", routes.GetPaymentsByUser)
	// app.Get("/api/payments/product/:id", routes.GetPaymentsByProduct)

	// POST routes
	// =================
	app.Post("/api/routess", routes.CreateUser)
	app.Post("/api/products", routes.CreateProduct)
	app.Post("/api/payments", routes.CreatePayment)

	// PUT routes
	// =================
	app.Put("/api/routess/:id", routes.UpdateUser)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Put("/api/payments/:id", routes.UpdatePayment)

	// DELETE routes
	// =================
	app.Delete("/api/routess/:id", routes.DeleteUser)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	app.Delete("/api/payments/:id", routes.DeletePayment)

}

func main() {
	database.Connect()
	app := fiber.New()

	setupRoutes(app)
	app.Listen(":3000")
}
