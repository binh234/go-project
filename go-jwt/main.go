package main

import (
	"jwt-demo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.GetInstance()
	initializers.SyncWithDB()
}

func authorize(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "authorize",
	})
}

func main() {
	r := gin.Default()

	r.GET("/auth", authorize)

	r.Run() // Listen and serve on port 8080 by default
}
