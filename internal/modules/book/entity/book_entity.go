package entity

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Title         string    `json:"title" validate:"required,min=2"`
	Author        string    `json:"author" validate:"required,min=2"`
	Genre         string    `json:"genre" validate:"required,min=2"`
	PublishedYear int       `json:"published_year" validate:"required,min=1900,max=2025"`
	Price         int       `json:"price" validate:"required,min=0"`	
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
