package controllers

import (
	"investo-rest-api/database"
	"investo-rest-api/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Nama:     data["nama"],
		Email:    data["email"],
		Password: data["password"],
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
