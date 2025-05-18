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

func (r *orderPostgres) Update(order *entity.Order) error {
	tx := r.db.Begin()

	// Delete existing items
	if err := tx.Where("order_id = ?", order.ID).Delete(&entity.OrderItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Save updated items
	for _, item := range order.Items {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update order total
	if err := tx.Model(order).Updates(map[string]interface{}{
		"total_price": order.TotalPrice,
		"updated_at":  order.UpdatedAt,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *orderPostgres) Delete(orderID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Delete items first due to foreign key constraint
		if err := tx.Where("order_id = ?", orderID).Delete(&entity.OrderItem{}).Error; err != nil {
			return err
		}
		// Then delete the order
		return tx.Where("id = ?", orderID).Delete(&entity.Order{}).Error
	})
}

func (r *orderPostgres) GetByID(orderID string) (*entity.Order, error) {
	var order entity.Order
	if err := r.db.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}