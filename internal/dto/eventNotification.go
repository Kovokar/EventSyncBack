package dto

import (
	"time"
)

// CreateEventNotificationRequest representa os dados de entrada para criação de notificação de evento
type CreateEventNotificationRequest struct {
	EventID       uint      `json:"event_id" binding:"required"`
	CreatedBy     uint      `json:"created_by" binding:"required"`
	Title         string    `json:"title" binding:"required"`
	Message       string    `json:"message" binding:"required"`
	TargetSegment string    `json:"target_segment" binding:"required,oneof=all approved check_in specific"`
	CreatedAtDate time.Time `json:"created_at_date"`
	SendEmail     bool      `json:"send_email"`
}

// UpdateEventNotificationRequest representa os dados de entrada para atualização de notificação de evento
type UpdateEventNotificationRequest struct {
	EventID       *uint      `json:"event_id,omitempty"`
	CreatedBy     *uint      `json:"created_by,omitempty"`
	Title         *string    `json:"title,omitempty" binding:"omitempty"`
	Message       *string    `json:"message,omitempty" binding:"omitempty"`
	TargetSegment *string    `json:"target_segment,omitempty" binding:"omitempty,oneof=all approved check_in specific"`
	CreatedAtDate *time.Time `json:"created_at_date,omitempty"`
	SendEmail     *bool      `json:"send_email,omitempty"`
}

// EventNotificationResponse representa a resposta da notificação de evento
type EventNotificationResponse struct {
	ID            uint      `json:"id"`
	EventID       uint      `json:"event_id"`
	CreatedBy     uint      `json:"created_by"`
	Title         string    `json:"title"`
	Message       string    `json:"message"`
	TargetSegment string    `json:"target_segment"`
	CreatedAtDate time.Time `json:"created_at_date"`
	SendEmail     bool      `json:"send_email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
