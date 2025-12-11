package models

import "time"

// Registration represents a user's registration for an event
type Registration struct {
	BaseModel
	EventID          uint      `gorm:"not null;index" binding:"required"`
	UserID           uint      `gorm:"not null;index" binding:"required"`
	Status           string    `gorm:"type:varchar(30);not null;default:'pending';index" binding:"required,oneof=pending approved rejected canceled"`
	RegistrationDate time.Time `gorm:"not null"`
	ApprovalDate     *time.Time
	CancellationDate *time.Time
	RejectionReason  string `gorm:"type:text"`
	CardQRCode       string `gorm:"type:varchar(500);uniqueIndex"`
	CheckInCount     int    `gorm:"default:0"`

	// Relationships
	Event    Event     `gorm:"foreignKey:EventID"`
	User     User      `gorm:"foreignKey:UserID"`
	Payment  *Payment  `gorm:"foreignKey:RegistrationID"`
	CheckIns []CheckIn `gorm:"foreignKey:RegistrationID"`
	// IssuedCertificates []IssuedCertificate `gorm:"foreignKey:RegistrationID"`
}

func (Registration) TableName() string { return "registrations" }
