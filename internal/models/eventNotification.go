package models

import "time"

// EventNotification represents a notification created by the organizer
type EventNotification struct {
	BaseModel
	EventID       uint      `gorm:"not null;index" binding:"required"`
	CreatedBy     uint      `gorm:"not null" binding:"required"`
	Title         string    `gorm:"type:varchar(255);not null" binding:"required"`
	Message       string    `gorm:"type:text;not null" binding:"required"`
	TargetSegment string    `gorm:"type:varchar(50);not null" binding:"required,oneof=all approved check_in specific"`
	CreatedAtDate time.Time `gorm:"not null"`
	SendEmail     bool      `gorm:"default:false"`

	// Relationships
	Event         Event              `gorm:"foreignKey:EventID"`
	Creator       User               `gorm:"foreignKey:CreatedBy"`
	Notifications []UserNotification `gorm:"foreignKey:EventNotificationID"`
}

func (EventNotification) TableName() string { return "event_notifications" }
