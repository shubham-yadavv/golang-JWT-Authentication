package config

import "github.com/shubham-yadavv/golang-JWT-Authentication/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
