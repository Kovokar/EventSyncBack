package models

import "time"

// Payment represents the payment of a registration
type Payment struct {
	BaseModel
	RegistrationID   uint   `gorm:"not null;uniqueIndex" binding:"required"`
	Amount           uint   `gorm:"type:decimal(10,2);not null" binding:"required"`
	Currency         string `gorm:"type:varchar(10);not null;default:'BRL'" binding:"required"`
	Method           string `gorm:"type:varchar(30);not null" binding:"required"`
	PaymentStatus    string `gorm:"type:varchar(30);not null;default:'pending';index" binding:"required,oneof=pending confirmed refunded"`
	PaymentDate      *time.Time
	ConfirmationDate *time.Time
	ReceiptURL       string `gorm:"type:varchar(500)"`
	Notes            string `gorm:"type:text"`

	// Relationship
	Registration Registration `gorm:"foreignKey:RegistrationID"`
}
