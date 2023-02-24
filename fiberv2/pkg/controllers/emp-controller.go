package controllers

import (
	"fibv2/pkg/database"
	"fibv2/pkg/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mg database.MongoInstance

func init() {
	database.Connect()
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
	id := c.Params("id")
	collection := mg.DB.Collection("employees")
	employeeID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	result := collection.FindOne(c.Context(), query)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString(result.Err().Error())
	}
	employee := new(models.Employee)
	result.Decode(employee)

	return c.JSON(employee)
}

func CreateNewEmployee(c *fiber.Ctx) error {
	collection := mg.DB.Collection("employees")
	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	employee.ID = ""
	insertionResult, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)
	createdEmployee := &models.Employee{}
	createdRecord.Decode(createdEmployee)

	return c.JSON(createdEmployee)
}

func UpdateEmployeeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	collection := mg.DB.Collection("employees")
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(404).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	updateOperators := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: employee.Name},
		{Key: "age", Value: employee.Age},
		{Key: "salary", Value: employee.Salary}}}}
	err = collection.FindOneAndUpdate(c.Context(), query, updateOperators).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employee)
}

func DeleteEmployeeByID(c *fiber.Ctx) error {
	collection := mg.DB.Collection("employees")
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := collection.DeleteOne(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
