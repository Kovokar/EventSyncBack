package dto

import (
	"time"

	"socialVoleiAPI/internal/models"
)

// CreateUserRequest representa os dados de entrada para criação de usuário
type CreateUserRequest struct {
	Name                string            `json:"name" binding:"required,min=3,max=100"`
	Email               string            `json:"email" binding:"required,email"`
	Password            string            `json:"password" binding:"required"`
	Birthdate           time.Time         `json:"birthdate" binding:"required"`
	Phone               string            `json:"phone" binding:"required,min=9,max=20"`
	Gender              models.GenderType `json:"gender" binding:"omitempty,oneof=male female other"`
	Photo               string            `json:"photo"`
	VisibleInPublicList bool              `json:"visible_in_public_list"`
}

// UpdateUserRequest representa os dados de entrada para atualização de usuário
type UpdateUserRequest struct {
	Name                *string            `json:"name,omitempty" binding:"omitempty,min=3,max=100"`
	Email               *string            `json:"email,omitempty" binding:"omitempty,email"`
	Password            *string            `json:"password,omitempty"`
	Birthdate           *time.Time         `json:"birthdate,omitempty"`
	Phone               *string            `json:"phone,omitempty" binding:"omitempty,min=9,max=20"`
	Gender              *models.GenderType `json:"gender,omitempty" binding:"omitempty,oneof=male female other"`
	Photo               *string            `json:"photo,omitempty"`
	VisibleInPublicList *bool              `json:"visible_in_public_list,omitempty"`
}

// UserResponse representa a resposta sem dados sensíveis
type UserResponse struct {
	ID                  uint              `json:"id"`
	Name                string            `json:"name"`
	Birthdate           time.Time         `json:"birthdate"`
	Email               string            `json:"email"`
	Phone               string            `json:"phone"`
	Gender              models.GenderType `json:"gender"`
	Photo               string            `json:"photo"`
	VisibleInPublicList bool              `json:"visible_in_public_list"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
}
