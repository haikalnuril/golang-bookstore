package postgres

import (
	"bookstore/internal/modules/order/entity"

	"gorm.io/gorm"
)

type orderPostgres struct {
	db *gorm.DB
}

func NewOrderPostgres(db *gorm.DB) *orderPostgres {
	return &orderPostgres{db: db}
}

func (r *orderPostgres) Create(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *orderPostgres) GetAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}
