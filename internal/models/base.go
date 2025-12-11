package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel defines common fields shared across all entities, such as ID,
// creation/update timestamps, soft delete support, and information about
// who created, updated, or deleted the record.

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedBy *uint          `json:"created_by,omitempty"`
	UpdatedBy *uint          `json:"updated_by,omitempty"`
	DeletedBy *uint          `json:"deleted_by,omitempty"`
}

func (BaseModel) TableName() string { return "base_model" }
