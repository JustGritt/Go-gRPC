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

/*
	func login(c *fiber.Ctx) error {
		user := c.FormValue("user")
		pass := c.FormValue("pass")

		// Throws Unauthorized error
		if user != "john" || pass != "doe" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"name":  "John Doe",
			"admin": true,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}
*/
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
	app.Get("/api/users", restricted, routes.GetUsers)
	app.Get("/api/users/:id", restricted, routes.GetUser)
	app.Get("/api/products", restricted, routes.GetProducts)
	app.Get("/api/products/:id", restricted, routes.GetProduct)
	app.Get("/api/payments", restricted, routes.GetPayments)
	app.Get("/api/payments/:id", restricted, routes.GetPayment)
	// app.Get("/api/payments/user/:id", routes.GetPaymentsByUser)
	// app.Get("/api/payments/product/:id", routes.GetPaymentsByProduct)

	// POST routes
	// =================
	app.Post("/api/users", restricted, routes.CreateUser)
	app.Post("/api/products", restricted, routes.CreateProduct)
	app.Post("/api/payments", restricted, routes.CreatePayment)

	// PUT routes
	// =================
	app.Put("/api/users/:id", restricted, routes.UpdateUser)
	app.Put("/api/products/:id", restricted, routes.UpdateProduct)
	app.Put("/api/payments/:id", restricted, routes.UpdatePayment)

	// DELETE routes
	// =================
	app.Delete("/api/users/:id", restricted, routes.DeleteUser)
	app.Delete("/api/products/:id", restricted, routes.DeleteProduct)
	app.Delete("/api/payments/:id", restricted, routes.DeletePayment)

}

func main() {
	database.Connect()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
