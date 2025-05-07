package postgres

import (
	"bookstore/internal/modules/book/entity"

	"gorm.io/gorm"
)

type bookPostgres struct {
	db *gorm.DB
}

func NewBookPostgres(db *gorm.DB) *bookPostgres {
	return &bookPostgres{db: db}
}

func (r *bookPostgres) Create(book *entity.Book) error {
	return r.db.Create(book).Error // r stand for repository
}

func (r *bookPostgres) GetAll() ([]entity.Book, error) {
	books := []entity.Book{}
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
