package models

import "time"

// Message represents a direct message exchanged between users
type Message struct {
	BaseModel
	SenderID      uint      `gorm:"not null;index" binding:"required"`
	ReceiverID    uint      `gorm:"not null;index" binding:"required"`
	Title         string    `gorm:"type:varchar(255);not null" binding:"required"`
	Body          string    `gorm:"type:text;not null" binding:"required"`
	AttachmentURL string    `gorm:"type:varchar(500)"`
	Read          bool      `gorm:"default:false"`
	SentAt        time.Time `gorm:"not null"`
	ReadAt        *time.Time

	// Relationships
	Sender   User `gorm:"foreignKey:SenderID"`
	Receiver User `gorm:"foreignKey:ReceiverID"`
}

func (Message) TableName() string { return "messages" }
