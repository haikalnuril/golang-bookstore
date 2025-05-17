package repository

import "bookstore/internal/modules/order/entity"

type OrderRepository interface {
	Create(order *entity.Order) error
	GetAll() ([]entity.Order, error)
}
