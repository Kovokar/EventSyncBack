package dto

import (
	"time"
)

// CreateReviewRequest representa os dados de entrada para criação de avaliação
type CreateReviewRequest struct {
	EventID    uint      `json:"event_id" binding:"required"`
	UserID     uint      `json:"user_id" binding:"required"`
	Rating     int       `json:"rating" binding:"required,min=1,max=5"`
	Comment    string    `json:"comment"`
	ReviewDate time.Time `json:"review_date"`
}

// UpdateReviewRequest representa os dados de entrada para atualização de avaliação
type UpdateReviewRequest struct {
	EventID    *uint      `json:"event_id,omitempty"`
	UserID     *uint      `json:"user_id,omitempty"`
	Rating     *int       `json:"rating,omitempty" binding:"omitempty,min=1,max=5"`
	Comment    *string    `json:"comment,omitempty"`
	ReviewDate *time.Time `json:"review_date,omitempty"`
}

// ReviewResponse representa a resposta da avaliação
type ReviewResponse struct {
	ID         uint      `json:"id"`
	EventID    uint      `json:"event_id"`
	UserID     uint      `json:"user_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	ReviewDate time.Time `json:"review_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
