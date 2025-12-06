package initializers

import "jwt_auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
