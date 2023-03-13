package initializers

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func connect() {
	dbName := os.Getenv("DB_NAME")
	connectDB, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	log.Println("Connected to database")
	db = connectDB
}

func GetInstance() *gorm.DB {
	once.Do(connect)
	return db
}
