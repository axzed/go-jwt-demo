package initializers

import (
	"github.com/axzed/go-jwt-demo/models"
)

func SyncDataBase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to sync database!")
	}
}
