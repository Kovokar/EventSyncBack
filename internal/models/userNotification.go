package models

import "time"

// UserNotification represents a notification sent to a specific user
type UserNotification struct {
	BaseModel
	EventNotificationID uint      `gorm:"not null;index" binding:"required"`
	UserID              uint      `gorm:"not null;index" binding:"required"`
	Read                bool      `gorm:"default:false"`
	SentAt              time.Time `gorm:"not null"`
	ReadAt              *time.Time

	// Relationships
	EventNotification EventNotification `gorm:"foreignKey:EventNotificationID"`
	User              User              `gorm:"foreignKey:UserID"`
}

func (UserNotification) TableName() string { return "user_notifications" }
