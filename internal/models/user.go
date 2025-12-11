package models

import (
	"time"
)

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
	Other  GenderType = "other"
)

type User struct {
	BaseModel
	Name                string     `gorm:"size:100;not null" binding:"required,min=3,max=100"`
	Email               string     `gorm:"size:100;uniqueIndex;not null" binding:"required,email"`
	Password            string     `gorm:"not null" binding:"required"`
	Birthdate           time.Time  `gorm:"not null" binding:"required"`
	Phone               string     `gorm:"size:50" binding:"required,min=9,max=20"`
	Gender              GenderType `gorm:"type:text;default:'other'" binding:"oneof=male female other"`
	Photo               string     `gorm:"default:'null'"`
	VisibleInPublicList bool       `gorm:"default:true"`

	// Relationships
	OrganizedEvents        []Event            `gorm:"foreignKey:OrganizerID"`
	Registrations          []Registration     `gorm:"foreignKey:UserID"`
	SentFriendRequests     []Friendship       `gorm:"foreignKey:RequesterUserID"`
	ReceivedFriendRequests []Friendship       `gorm:"foreignKey:RecipientUserID"`
	SentMessages           []Message          `gorm:"foreignKey:SenderID"`
	ReceivedMessages       []Message          `gorm:"foreignKey:ReceiverID"`
	Reviews                []Review           `gorm:"foreignKey:UserID"`
	Notifications          []UserNotification `gorm:"foreignKey:UserID"`
}

func (User) TableName() string { return "users" }
