package database

import (
	"context"
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg MongoInstance

const dbName = "hrms"
const mongoURI = "mongodb+srv://<username>:<password>@cluster0.9fkxa.mongodb.net/?retryWrites=true&w=majority"

var serverAPIOptions = options.ServerAPI(options.ServerAPIVersion1)

func Connect() error {
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		DB:     db,
	}
	return nil
}

func GetMongoInstance() MongoInstance {
	return mg
}
