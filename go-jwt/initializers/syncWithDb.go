package initializers

import "jwt-demo/models"

func SyncWithDB() {
	db.AutoMigrate(&models.User{})
}
