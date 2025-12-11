package dto

import (
	"time"
)

// CreateMessageRequest representa os dados de entrada para criação de mensagem
type CreateMessageRequest struct {
	SenderID      uint       `json:"sender_id" binding:"required"`
	ReceiverID    uint       `json:"receiver_id" binding:"required"`
	Title         string     `json:"title" binding:"required"`
	Body          string     `json:"body" binding:"required"`
	AttachmentURL string     `json:"attachment_url"`
	Read          bool       `json:"read"`
	SentAt        time.Time  `json:"sent_at"`
	ReadAt        *time.Time `json:"read_at,omitempty"`
}

// UpdateMessageRequest representa os dados de entrada para atualização de mensagem
type UpdateMessageRequest struct {
	SenderID      *uint      `json:"sender_id,omitempty"`
	ReceiverID    *uint      `json:"receiver_id,omitempty"`
	Title         *string    `json:"title,omitempty" binding:"omitempty"`
	Body          *string    `json:"body,omitempty" binding:"omitempty"`
	AttachmentURL *string    `json:"attachment_url,omitempty"`
	Read          *bool      `json:"read,omitempty"`
	SentAt        *time.Time `json:"sent_at,omitempty"`
	ReadAt        *time.Time `json:"read_at,omitempty"`
}

// MessageResponse representa a resposta da mensagem
type MessageResponse struct {
	ID            uint       `json:"id"`
	SenderID      uint       `json:"sender_id"`
	ReceiverID    uint       `json:"receiver_id"`
	Title         string     `json:"title"`
	Body          string     `json:"body"`
	AttachmentURL string     `json:"attachment_url"`
	Read          bool       `json:"read"`
	SentAt        time.Time  `json:"sent_at"`
	ReadAt        *time.Time `json:"read_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
