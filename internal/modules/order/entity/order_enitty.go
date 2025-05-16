package entity

import	(
		"time"
		"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID    `gorm:"type:uuid" json:"user_id"`
	Items      []OrderItem  `gorm:"foreignKey:OrderID" json:"items"`
	TotalPrice int          `json:"total_price"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}