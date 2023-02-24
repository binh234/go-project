package database

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg *MongoInstance
var once sync.Once

const dbName = "graphql"

var mongoURI string
var serverAPIOptions = options.ServerAPI(options.ServerAPIVersion1)

func init() {
	// Set up the configuration file
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()

	// Get the secret key from the configuration file
	mongoURI = viper.GetString("MONGO_URI")
}

// Connect to MongoDB
func connect() *MongoInstance {
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(dbName)
	return &MongoInstance{
		Client: client,
		DB:     db,
	}
}

// Singleton pattern
func GetMongoInstance() *MongoInstance {
	once.Do(func() {
		mg = connect()
	})
	return mg
}
