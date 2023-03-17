package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

// var rdb *redis.Client
// var once sync.Once

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}

// func GetInstance(dbNo int) *redis.Client {
// 	once.Do(func() { createClient(dbNo) })
// 	return rdb
// }
