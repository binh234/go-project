package main

import (
	"fibv2/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/employee")
	api.Get("/", controllers.GetAllEmployees)
	api.Get("/:id", controllers.GetEmployeeByID)
	api.Post("/employee", controllers.CreateNewEmployee)
	api.Put("/:id", controllers.UpdateEmployeeByID)
	api.Delete("/:id", controllers.DeleteEmployeeByID)
}
