package models

import (
	"time"
	// "gorm.io/gorm"
)

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
	Other  GenderType = "other"
)

type User struct {
	BaseModel
	Name      string     `gorm:"size:100;not null" binding:"required,min=3,max=100"`
	Email     string     `gorm:"size:100;uniqueIndex;not null" binding:"required,email"`
	Password  string     `gorm:"not null" binding:"required"`
	Birthdate time.Time  `gorm:"not null" binding:"required"`
	Phone     string     `gorm:"size:50" binding:"required,min=9,max=20"`
	Gender    GenderType `gorm:"type:text;default:'other'" binding:"oneof=male female other"`
	Photo     string     `gorm:"default:'null'" json:"photo"`
	// VisivelListaPublica  bool      `gorm:"default:true" json:"visivel_lista_publica"`
}

// UsuarioResponse representa a resposta sem dados sensíveis
type UserResponse struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	Birthdate time.Time  `json:"birthdate" `
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Gender    GenderType `json:"gender"`
	Photo     string     `json:"photo"`
}
