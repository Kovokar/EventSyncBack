package models

import "time"

// EventNotification represents a notification created by the organizer
type EventNotification struct {
	BaseModel
	EventID       uint      `gorm:"not null;index" json:"event_id" binding:"required"`
	CreatedBy     uint      `gorm:"not null" json:"created_by" binding:"required"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	Message       string    `gorm:"type:text;not null" json:"message" binding:"required"`
	TargetSegment string    `gorm:"type:varchar(50);not null" json:"target_segment" binding:"required,oneof=all approved check_in specific"`
	CreatedAtDate time.Time `gorm:"not null" json:"created_at_date"`
	SendEmail     bool      `gorm:"default:false" json:"send_email"`

	// Relationships
	Event         Event              `gorm:"foreignKey:EventID" json:"event,omitempty"`
	Creator       User               `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	Notifications []UserNotification `gorm:"foreignKey:EventNotificationID" json:"notifications,omitempty"`
}

func (EventNotification) TableName() string { return "nevent_notifications" }
