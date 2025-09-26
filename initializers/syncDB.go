package initializers

import "github.com/Pkra99/go-jwt-auth/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}