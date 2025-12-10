package migrations

import (
	"socialVoleiAPI/internal/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
