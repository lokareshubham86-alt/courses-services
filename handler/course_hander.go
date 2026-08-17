package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"courses-services/services"
)

func GetCourses(c *fiber.Ctx) error {
	log.Println("The control is coming in the handler")
	courses, err := services.GetCourses()
	if err != nil {
		log.Println("The error is : ", err)
		c.Status(http.StatusBadGateway).JSON(err.Error())
	}
	return c.JSON(courses)
}
