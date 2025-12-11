package models

import "time"

type Event struct {
	BaseModel
	OrganizerID           uint      `gorm:"not null;index" binding:"required"`
	Title                 string    `gorm:"type:varchar(255);not null" binding:"required"`
	ShortDescription      string    `gorm:"type:text;not null" binding:"required"`
	FullDescription       string    `gorm:"type:text"`
	LocationType          string    `gorm:"type:varchar(20);not null" binding:"required,oneof=online in_person"`
	LocationValue         string    `gorm:"type:text;not null" binding:"required"`
	StartDate             time.Time `gorm:"not null;index" binding:"required"`
	EndDate               time.Time `gorm:"not null" binding:"required"`
	Capacity              *int
	EventType             string `gorm:"type:varchar(20);not null" binding:"required,oneof=free paid"`
	Price                 *uint  `gorm:"type:decimal(10,2)"`
	PaymentInstructions   string `gorm:"type:text"`
	PixKey                string `gorm:"type:varchar(255)"`
	AutomaticRegistration bool   `gorm:"default:false"`
	DisplayParticipants   bool   `gorm:"default:false"`
	BannerURL             string `gorm:"type:varchar(500)"`
	EventQRCode           string `gorm:"type:varchar(500)"`
	Status                string `gorm:"type:varchar(30);not null;default:'draft';index" binding:"required,oneof=draft open_for_registration in_progress finished"`
	RegistrationOpenDate  *time.Time
	RegistrationCloseDate *time.Time
	WorkloadHours         *int
	Category              string `gorm:"type:varchar(50)"`
	Highlighted           bool   `gorm:"default:false"`

	// Relationships
	Organizer     User           `gorm:"foreignKey:OrganizerID"`
	Registrations []Registration `gorm:"foreignKey:EventID"`
	CheckIns      []CheckIn      `gorm:"foreignKey:EventID"`
	Reviews       []Review       `gorm:"foreignKey:EventID"`
	// CertificateConfig *CertificateConfig  `gorm:"foreignKey:EventID"`
	Notifications []EventNotification `gorm:"foreignKey:EventID"`
}

func (Event) TableName() string { return "events" }
