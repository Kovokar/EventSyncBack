package models

import "time"

// Review represents the evaluation of an event by a participant
type Review struct {
	BaseModel
	EventID    uint      `gorm:"not null;index;uniqueIndex:idx_event_user" binding:"required"`
	UserID     uint      `gorm:"not null;index;uniqueIndex:idx_event_user" binding:"required"`
	Rating     int       `gorm:"not null" binding:"required,min=1,max=5"`
	Comment    string    `gorm:"type:text"`
	ReviewDate time.Time `gorm:"not null"`

	// Relationships
	Event Event `gorm:"foreignKey:EventID"`
	User  User  `gorm:"foreignKey:UserID"`
}

func (Review) TableName() string { return "reviews" }
