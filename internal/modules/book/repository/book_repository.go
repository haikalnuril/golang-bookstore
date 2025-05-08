package repository

import "bookstore/internal/modules/book/entity"

type BookRepository interface {
	Create(book *entity.Book) error
	GetAll() ([]entity.Book, error)
}
