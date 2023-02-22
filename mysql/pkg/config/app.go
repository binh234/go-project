package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// Connect to the database
	d, err := gorm.Open("mysql", "admin:admin@tcp(127.0.0.1:3336)/bookstore_sample?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
