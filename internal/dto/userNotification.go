package dto

import (
	"time"
)

// CreateUserNotificationRequest representa os dados de entrada para criação de notificação de usuário
type CreateUserNotificationRequest struct {
	EventNotificationID uint       `json:"event_notification_id" binding:"required"`
	UserID              uint       `json:"user_id" binding:"required"`
	Read                bool       `json:"read"`
	SentAt              time.Time  `json:"sent_at"`
	ReadAt              *time.Time `json:"read_at,omitempty"`
}

// UpdateUserNotificationRequest representa os dados de entrada para atualização de notificação de usuário
type UpdateUserNotificationRequest struct {
	EventNotificationID *uint      `json:"event_notification_id,omitempty"`
	UserID              *uint      `json:"user_id,omitempty"`
	Read                *bool      `json:"read,omitempty"`
	SentAt              *time.Time `json:"sent_at,omitempty"`
	ReadAt              *time.Time `json:"read_at,omitempty"`
}

// UserNotificationResponse representa a resposta da notificação de usuário
type UserNotificationResponse struct {
	ID                  uint       `json:"id"`
	EventNotificationID uint       `json:"event_notification_id"`
	UserID              uint       `json:"user_id"`
	Read                bool       `json:"read"`
	SentAt              time.Time  `json:"sent_at"`
	ReadAt              *time.Time `json:"read_at,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
