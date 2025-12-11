package dto

import (
	"time"
)

// CreateCheckInRequest representa os dados de entrada para criação de check-in
type CreateCheckInRequest struct {
	RegistrationID uint      `json:"registration_id" binding:"required"`
	EventID        uint      `json:"event_id" binding:"required"`
	DateTime       time.Time `json:"date_time" binding:"required"`
	Method         string    `json:"method" binding:"required,oneof=manual qrcode code"`
	PerformedBy    uint      `json:"performed_by" binding:"required"`
	Notes          string    `json:"notes"`
}

// UpdateCheckInRequest representa os dados de entrada para atualização de check-in
type UpdateCheckInRequest struct {
	RegistrationID *uint      `json:"registration_id,omitempty"`
	EventID        *uint      `json:"event_id,omitempty"`
	DateTime       *time.Time `json:"date_time,omitempty" binding:"omitempty"`
	Method         *string    `json:"method,omitempty" binding:"omitempty,oneof=manual qrcode code"`
	PerformedBy    *uint      `json:"performed_by,omitempty"`
	Notes          *string    `json:"notes,omitempty"`
}

// CheckInResponse representa a resposta do check-in
type CheckInResponse struct {
	ID             uint      `json:"id"`
	RegistrationID uint      `json:"registration_id"`
	EventID        uint      `json:"event_id"`
	DateTime       time.Time `json:"date_time"`
	Method         string    `json:"method"`
	PerformedBy    uint      `json:"performed_by"`
	Notes          string    `json:"notes"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
