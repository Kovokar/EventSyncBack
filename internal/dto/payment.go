package dto

import (
	"time"
)

// CreatePaymentRequest representa os dados de entrada para criação de pagamento
type CreatePaymentRequest struct {
	RegistrationID   uint       `json:"registration_id" binding:"required"`
	Amount           uint       `json:"amount" binding:"required"`
	Currency         string     `json:"currency" binding:"required"`
	Method           string     `json:"method" binding:"required"`
	PaymentStatus    string     `json:"payment_status" binding:"required,oneof=pending confirmed refunded"`
	PaymentDate      *time.Time `json:"payment_date,omitempty"`
	ConfirmationDate *time.Time `json:"confirmation_date,omitempty"`
	ReceiptURL       string     `json:"receipt_url"`
	Notes            string     `json:"notes"`
}

// UpdatePaymentRequest representa os dados de entrada para atualização de pagamento
type UpdatePaymentRequest struct {
	RegistrationID   *uint      `json:"registration_id,omitempty"`
	Amount           *uint      `json:"amount,omitempty"`
	Currency         *string    `json:"currency,omitempty"`
	Method           *string    `json:"method,omitempty"`
	PaymentStatus    *string    `json:"payment_status,omitempty" binding:"omitempty,oneof=pending confirmed refunded"`
	PaymentDate      *time.Time `json:"payment_date,omitempty"`
	ConfirmationDate *time.Time `json:"confirmation_date,omitempty"`
	ReceiptURL       *string    `json:"receipt_url,omitempty"`
	Notes            *string    `json:"notes,omitempty"`
}

// PaymentResponse representa a resposta do pagamento
type PaymentResponse struct {
	ID               uint       `json:"id"`
	RegistrationID   uint       `json:"registration_id"`
	Amount           uint       `json:"amount"`
	Currency         string     `json:"currency"`
	Method           string     `json:"method"`
	PaymentStatus    string     `json:"payment_status"`
	PaymentDate      *time.Time `json:"payment_date,omitempty"`
	ConfirmationDate *time.Time `json:"confirmation_date,omitempty"`
	ReceiptURL       string     `json:"receipt_url"`
	Notes            string     `json:"notes"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
