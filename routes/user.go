package routes

import (
	"github.com/JustGritt/go-grpc/database"
	"github.com/JustGritt/go-grpc/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

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
