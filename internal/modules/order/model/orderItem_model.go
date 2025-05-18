package model

import (
	"github.com/google/uuid"
)

type CreateOrderItemInput struct {
	BookID   uuid.UUID `json:"book_id"`
	Quantity int       `json:"quantity"`
}

type UpdateOrderItemRequest struct {
	BookID   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}

type OrderItemResponse struct {
	ID        string `json:"id"`
	BookID    string `json:"book_id"`
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unit_price"`
}