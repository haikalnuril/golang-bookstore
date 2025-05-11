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

func (r *bookPostgres) GetByID(id string) (*entity.Book, error) {
	book := &entity.Book{}
	err := r.db.First(book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *bookPostgres) Update(id string, updates map[string]interface{}) (*entity.Book, error) {
	err := r.db.Model(&entity.Book{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	var updated entity.Book
	if err := r.db.First(&updated, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *bookPostgres) Delete(id string) error {
	return r.db.Delete(&entity.Book{}, "id = ?", id).Error
}
