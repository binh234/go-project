package main

import (
	"fibv2/pkg/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/employee")
	api.Get("/", controllers.GetAllEmployees)
	api.Get("/:id", controllers.GetEmployeeByID)
	api.Post("/", controllers.CreateNewEmployee)
	api.Put("/:id", controllers.UpdateEmployeeByID)
	api.Delete("/:id", controllers.DeleteEmployeeByID)

	log.Fatal(app.Listen(":8080"))
}
