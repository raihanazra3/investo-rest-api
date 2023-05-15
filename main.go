package main

import (
	"investo-rest-api/database"
	"investo-rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":1337")
}
