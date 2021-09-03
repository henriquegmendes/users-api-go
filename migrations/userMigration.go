package migrations

import (
	"gorm.io/gorm"
	"henrique.mendes/users-api/models"
)

func CreateTables(database *gorm.DB) {
	database.AutoMigrate(&models.User{})
}
