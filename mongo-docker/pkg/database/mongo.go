package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg *MongoInstance
var once sync.Once

const dbName = "hrms"

var mongoURI string = os.Getenv("CONFIG_MONGO_URI")
var serverAPIOptions = options.ServerAPI(options.ServerAPIVersion1)

func connect() {
	fmt.Println(mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	db := client.Database(dbName)
	mg = &MongoInstance{
		Client: client,
		DB:     db,
	}
	log.Println("Connect succesfully")
}

func GetMongoInstance() *MongoInstance {
	once.Do(connect)
	return mg
}
