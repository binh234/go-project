package controllers

import (
	"fibv2/pkg/database"
	"fibv2/pkg/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var mg database.MongoInstance

func init() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	mg = database.GetMongoInstance()
}

func GetAllEmployees(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := mg.DB.Collection("employees").Find(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []models.Employee = make([]models.Employee, 0)
	err = cursor.All(c.Context(), &employees)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}

func GetEmployeeByID(c *fiber.Ctx) error {
	return nil
}

func CreateNewEmployee(c *fiber.Ctx) error {
	return nil
}

func UpdateEmployeeByID(c *fiber.Ctx) error {
	return nil
}

func DeleteEmployeeByID(c *fiber.Ctx) error {
	return nil
}
