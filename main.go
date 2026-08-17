package main

import (
	"github.com/gofiber/fiber/v2"

	"courses-services/database"
	"courses-services/routes"
)

func main() {
	database.ConnectToDatabase()
	database.CreateTable()

	app := fiber.New()
	//Routes set up
	routes.SetUpRoutes(app)
	app.Listen(":8000")
}
