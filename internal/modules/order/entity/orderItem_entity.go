package entity

import (
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid" json:"order_id"`
	BookID    uuid.UUID `gorm:"type:uuid" json:"book_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice int       `json:"unit_price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
