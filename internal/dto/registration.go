package dto

import (
	"time"
)

// CreateRegistrationRequest representa os dados de entrada para criação de inscrição
type CreateRegistrationRequest struct {
	EventID          uint       `json:"event_id" binding:"required"`
	UserID           uint       `json:"user_id" binding:"required"`
	Status           string     `json:"status" binding:"required,oneof=pending approved rejected canceled"`
	RegistrationDate time.Time  `json:"registration_date"`
	ApprovalDate     *time.Time `json:"approval_date,omitempty"`
	CancellationDate *time.Time `json:"cancellation_date,omitempty"`
	RejectionReason  string     `json:"rejection_reason"`
	CardQRCode       string     `json:"card_qr_code"`
	CheckInCount     int        `json:"check_in_count"`
}

// UpdateRegistrationRequest representa os dados de entrada para atualização de inscrição
type UpdateRegistrationRequest struct {
	EventID          *uint      `json:"event_id,omitempty"`
	UserID           *uint      `json:"user_id,omitempty"`
	Status           *string    `json:"status,omitempty" binding:"omitempty,oneof=pending approved rejected canceled"`
	RegistrationDate *time.Time `json:"registration_date,omitempty"`
	ApprovalDate     *time.Time `json:"approval_date,omitempty"`
	CancellationDate *time.Time `json:"cancellation_date,omitempty"`
	RejectionReason  *string    `json:"rejection_reason,omitempty"`
	CardQRCode       *string    `json:"card_qr_code,omitempty"`
	CheckInCount     *int       `json:"check_in_count,omitempty"`
}

// RegistrationResponse representa a resposta da inscrição
type RegistrationResponse struct {
	ID               uint       `json:"id"`
	EventID          uint       `json:"event_id"`
	UserID           uint       `json:"user_id"`
	Status           string     `json:"status"`
	RegistrationDate time.Time  `json:"registration_date"`
	ApprovalDate     *time.Time `json:"approval_date,omitempty"`
	CancellationDate *time.Time `json:"cancellation_date,omitempty"`
	RejectionReason  string     `json:"rejection_reason"`
	CardQRCode       string     `json:"card_qr_code"`
	CheckInCount     int        `json:"check_in_count"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
