package dto

import (
	"time"
)

// CreateEventRequest representa os dados de entrada para criação de evento
type CreateEventRequest struct {
	OrganizerID           uint       `json:"organizer_id" binding:"required"`
	Title                 string     `json:"title" binding:"required"`
	ShortDescription      string     `json:"short_description" binding:"required"`
	FullDescription       string     `json:"full_description"`
	LocationType          string     `json:"location_type" binding:"required,oneof=online in_person"`
	LocationValue         string     `json:"location_value" binding:"required"`
	StartDate             time.Time  `json:"start_date" binding:"required"`
	EndDate               time.Time  `json:"end_date" binding:"required"`
	Capacity              *int       `json:"capacity,omitempty"`
	EventType             string     `json:"event_type" binding:"required,oneof=free paid"`
	Price                 *uint      `json:"price,omitempty"`
	PaymentInstructions   string     `json:"payment_instructions"`
	PixKey                string     `json:"pix_key"`
	AutomaticRegistration bool       `json:"automatic_registration"`
	DisplayParticipants   bool       `json:"display_participants"`
	BannerURL             string     `json:"banner_url"`
	EventQRCode           string     `json:"event_qr_code"`
	Status                string     `json:"status" binding:"required,oneof=draft open_for_registration in_progress finished"`
	RegistrationOpenDate  *time.Time `json:"registration_open_date,omitempty"`
	RegistrationCloseDate *time.Time `json:"registration_close_date,omitempty"`
	WorkloadHours         *int       `json:"workload_hours,omitempty"`
	Category              string     `json:"category"`
	Highlighted           bool       `json:"highlighted"`
}

// UpdateEventRequest representa os dados de entrada para atualização de evento
type UpdateEventRequest struct {
	OrganizerID           *uint      `json:"organizer_id,omitempty"`
	Title                 *string    `json:"title,omitempty" binding:"omitempty"`
	ShortDescription      *string    `json:"short_description,omitempty" binding:"omitempty"`
	FullDescription       *string    `json:"full_description,omitempty"`
	LocationType          *string    `json:"location_type,omitempty" binding:"omitempty,oneof=online in_person"`
	LocationValue         *string    `json:"location_value,omitempty" binding:"omitempty"`
	StartDate             *time.Time `json:"start_date,omitempty" binding:"omitempty"`
	EndDate               *time.Time `json:"end_date,omitempty" binding:"omitempty"`
	Capacity              *int       `json:"capacity,omitempty"`
	EventType             *string    `json:"event_type,omitempty" binding:"omitempty,oneof=free paid"`
	Price                 *uint      `json:"price,omitempty"`
	PaymentInstructions   *string    `json:"payment_instructions,omitempty"`
	PixKey                *string    `json:"pix_key,omitempty"`
	AutomaticRegistration *bool      `json:"automatic_registration,omitempty"`
	DisplayParticipants   *bool      `json:"display_participants,omitempty"`
	BannerURL             *string    `json:"banner_url,omitempty"`
	EventQRCode           *string    `json:"event_qr_code,omitempty"`
	Status                *string    `json:"status,omitempty" binding:"omitempty,oneof=draft open_for_registration in_progress finished"`
	RegistrationOpenDate  *time.Time `json:"registration_open_date,omitempty"`
	RegistrationCloseDate *time.Time `json:"registration_close_date,omitempty"`
	WorkloadHours         *int       `json:"workload_hours,omitempty"`
	Category              *string    `json:"category,omitempty"`
	Highlighted           *bool      `json:"highlighted,omitempty"`
}

// EventResponse representa a resposta do evento
type EventResponse struct {
	ID                    uint       `json:"id"`
	OrganizerID           uint       `json:"organizer_id"`
	Title                 string     `json:"title"`
	ShortDescription      string     `json:"short_description"`
	FullDescription       string     `json:"full_description"`
	LocationType          string     `json:"location_type"`
	LocationValue         string     `json:"location_value"`
	StartDate             time.Time  `json:"start_date"`
	EndDate               time.Time  `json:"end_date"`
	Capacity              *int       `json:"capacity,omitempty"`
	EventType             string     `json:"event_type"`
	Price                 *uint      `json:"price,omitempty"`
	PaymentInstructions   string     `json:"payment_instructions"`
	PixKey                string     `json:"pix_key"`
	AutomaticRegistration bool       `json:"automatic_registration"`
	DisplayParticipants   bool       `json:"display_participants"`
	BannerURL             string     `json:"banner_url"`
	EventQRCode           string     `json:"event_qr_code"`
	Status                string     `json:"status"`
	RegistrationOpenDate  *time.Time `json:"registration_open_date,omitempty"`
	RegistrationCloseDate *time.Time `json:"registration_close_date,omitempty"`
	WorkloadHours         *int       `json:"workload_hours,omitempty"`
	Category              string     `json:"category"`
	Highlighted           bool       `json:"highlighted"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
}
