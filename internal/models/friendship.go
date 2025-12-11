package models

import "time"

// Friendship represents a friendship relationship between users
type Friendship struct {
	BaseModel
	RequesterUserID uint      `gorm:"not null;index" binding:"required"`
	RecipientUserID uint      `gorm:"not null;index" binding:"required"`
	SourceEventID   uint      `gorm:"not null;index" binding:"required"`
	Status          string    `gorm:"type:varchar(30);not null;default:'pending';index" binding:"required,oneof=pending accepted rejected undone"`
	RequestDate     time.Time `gorm:"not null"`
	ResponseDate    *time.Time

	// Relationships
	RequesterUser User  `gorm:"foreignKey:RequesterUserID"`
	RecipientUser User  `gorm:"foreignKey:RecipientUserID"`
	SourceEvent   Event `gorm:"foreignKey:SourceEventID"`
}

func (Friendship) TableName() string { return "Friendships" }
