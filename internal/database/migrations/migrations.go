package migrations

import (
	"socialVoleiAPI/internal/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.BaseModel{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Event{})
	db.AutoMigrate(models.CheckIn{})
	db.AutoMigrate(models.EventNotification{})
	db.AutoMigrate(models.Friendship{})
	db.AutoMigrate(models.Message{})
	db.AutoMigrate(models.Payment{})
	db.AutoMigrate(models.Registration{})
	db.AutoMigrate(models.Review{})
	db.AutoMigrate(models.UserNotification{})
}
