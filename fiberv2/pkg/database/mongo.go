package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg MongoInstance

const dbName = "hrms"

var mongoURI string

var serverAPIOptions = options.ServerAPI(options.ServerAPIVersion1)

func getConfig() {
	// Set up the configuration file
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()

	// Get the secret key from the configuration file
	mongoURI = viper.GetString("MONGO_URI")
}

func Connect() {
	getConfig()
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		DB:     db,
	}
}

func GetMongoInstance() MongoInstance {
	return mg
}
