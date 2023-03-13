package main

import (
	"jwt-demo/controllers"
	"jwt-demo/initializers"
	"jwt-demo/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Initializing...")
	initializers.LoadEnvVariables()
	initializers.GetInstance()
	initializers.SyncWithDB()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "homepage",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", home)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/logout", controllers.Logout)

	r.Run() // Listen and serve on port 8080 by default
}
