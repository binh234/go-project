package initializers

import (
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var once *sync.Once
var db *gorm.DB

func connect() {
	dbName := os.Getenv("DB_NAME")
	connectDB, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = connectDB
}

func GetInstance() *gorm.DB {
	once.Do(connect)
	return db
}
