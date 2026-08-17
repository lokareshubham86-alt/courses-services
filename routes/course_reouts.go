package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"courses-services/handler"
)

func SetUpRoutes(app *fiber.App) {
	log.Println("The control is coming in the routes")
	app.Get("/courses", handler.GetCourses)
	app.Get("/courses/:id", handler.GetCourseByID)
}
