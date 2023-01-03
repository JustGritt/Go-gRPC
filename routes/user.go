package routes

import (
	"github.com/JustGritt/go-grpc/database"
	"github.com/JustGritt/go-grpc/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func validToken(t *jwt.Token, id int) bool {
	n := id
	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	return uid == n
}

func CreateResponseUser(user models.User) User {
	return User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// encrypt password with bcrypt
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error encrypting password",
		})
	}

	user.Password = string(pass)

	database.Database.Db.Create(&user)

	return c.Status(201).JSON(CreateResponseUser(user))
}

func GetUserId(id int, user *models.User) error {
	database.Database.Db.Find(&user, id)
	if user.ID == 0 {
		return fiber.NewError(404, "User not found")
	}

	return nil
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)

	var responseUsers []User
	for _, user := range users {
		responseUsers = append(responseUsers, CreateResponseUser(user))
	}

	return c.Status(200).JSON(responseUsers)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.Database.Db.Find(&user, id)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(200).JSON(CreateResponseUser(user))
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = GetUserId(id, &user)

	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	// encrypt password with bcrypt
	pass, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), 14)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error encrypting password",
		})
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.Password = string(pass)

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = GetUserId(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Delete(&user)

	return c.Status(200).JSON(fiber.Map{
		"message": "User deleted",
	})
}
