package repository

import "bookstore/internal/modules/book/entity"

type BookRepository interface {
	Create(book *entity.Book) error
	GetAll() ([]entity.Book, error)
	GetByID(id string) (*entity.Book, error)
	Update(id string, book *entity.Book) (*entity.Book, error)
	Delete(id string) error
}