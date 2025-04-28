package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name" validate:"required,min=2"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
