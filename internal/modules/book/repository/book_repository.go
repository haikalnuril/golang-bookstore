package repository

import (
	"bookstore/internal/modules/book/entity"
	"bookstore/internal/modules/book/model"
)

type BookRepository interface {
	Create(book *entity.Book) error
	GetAll(req *model.SearchBookRequest) ([]entity.Book, int, error)
	GetByID(id string) (*entity.Book, error)
	Update(id string, updates map[string]interface{}) (*entity.Book, error)
	Delete(id string) error
}
