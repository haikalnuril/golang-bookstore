package repository

import "bookstore/internal/modules/order/entity"

type OrderRepository interface {
	Create(order *entity.Order) error
	GetAll() ([]entity.Order, error)
	Update(order *entity.Order) error
	Delete(orderID string) error
	GetByID(orderID string) (*entity.Order, error)
}
