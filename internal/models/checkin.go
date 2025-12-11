package models

import "time"

// CheckIn represents the attendance record for an event
type CheckIn struct {
	BaseModel
	RegistrationID uint      `gorm:"not null;index" binding:"required"`
	EventID        uint      `gorm:"not null;index" binding:"required"`
	DateTime       time.Time `gorm:"not null" binding:"required"`
	Method         string    `gorm:"type:varchar(30);not null" binding:"required,oneof=manual qrcode code"`
	PerformedBy    uint      `gorm:"not null" binding:"required"`
	Notes          string    `gorm:"type:text"`

	// Relationships
	Registration Registration `gorm:"foreignKey:RegistrationID"`
	Event        Event        `gorm:"foreignKey:EventID"`
	Organizer    User         `gorm:"foreignKey:PerformedBy"`
}

func (CheckIn) TableName() string { return "check_ins" }
