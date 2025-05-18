package model

import (
	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	UserID uuid.UUID              `json:"user_id"`
	Items  []CreateOrderItemInput `json:"items"` // <- use submodel type from orderItem_model.go
}

type UpdateOrderRequest struct {
	Items []UpdateOrderItemRequest `json:"items"`
}

type OrderResponse struct {
	ID         string              `json:"id"`
	UserID     string              `json:"user_id"`
	Items      []OrderItemResponse `json:"items"` // <- use submodel type from orderItem_model.go
	TotalPrice int                 `json:"total_price"`
	CreatedAt  string              `json:"created_at"`
	UpdatedAt  string              `json:"updated_at"`
}
