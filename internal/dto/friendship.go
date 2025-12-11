package dto

import (
	"time"
)

// CreateFriendshipRequest representa os dados de entrada para criação de amizade
type CreateFriendshipRequest struct {
	RequesterUserID uint       `json:"requester_user_id" binding:"required"`
	RecipientUserID uint       `json:"recipient_user_id" binding:"required"`
	SourceEventID   uint       `json:"source_event_id" binding:"required"`
	Status          string     `json:"status" binding:"required,oneof=pending accepted rejected undone"`
	RequestDate     time.Time  `json:"request_date"`
	ResponseDate    *time.Time `json:"response_date,omitempty"`
}

// UpdateFriendshipRequest representa os dados de entrada para atualização de amizade
type UpdateFriendshipRequest struct {
	RequesterUserID *uint      `json:"requester_user_id,omitempty"`
	RecipientUserID *uint      `json:"recipient_user_id,omitempty"`
	SourceEventID   *uint      `json:"source_event_id,omitempty"`
	Status          *string    `json:"status,omitempty" binding:"omitempty,oneof=pending accepted rejected undone"`
	RequestDate     *time.Time `json:"request_date,omitempty"`
	ResponseDate    *time.Time `json:"response_date,omitempty"`
}

// FriendshipResponse representa a resposta da amizade
type FriendshipResponse struct {
	ID              uint       `json:"id"`
	RequesterUserID uint       `json:"requester_user_id"`
	RecipientUserID uint       `json:"recipient_user_id"`
	SourceEventID   uint       `json:"source_event_id"`
	Status          string     `json:"status"`
	RequestDate     time.Time  `json:"request_date"`
	ResponseDate    *time.Time `json:"response_date,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
