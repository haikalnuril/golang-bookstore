package usecase

import (
	"bookstore/internal/modules/book/entity"
	"bookstore/internal/modules/book/model"
	"bookstore/internal/modules/book/repository"
	"time"
)

type BookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) *BookUsecase {
	return &BookUsecase{repo: repo}
}

func (b *BookUsecase) Create(req *model.BookRequest) (*model.BookResponse, error) {
	book := &entity.Book{
		Title:         req.Title,
		Author:        req.Author,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
		Price:         req.Price,
	}

	err := b.repo.Create(book)
	if err != nil {
		return nil, err
	}
	return &model.BookResponse{
		ID:            book.ID.String(),
		Title:         book.Title,
		Author:        book.Author,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
		Price:         book.Price,
		CreatedAt:     book.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     book.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (b *BookUsecase) GetAll() ([]model.BookResponse, error) {
	books, err := b.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var bookResponses []model.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, model.BookResponse{
			ID:            book.ID.String(),
			Title:         book.Title,
			Author:        book.Author,
			Genre:         book.Genre,
			PublishedYear: book.PublishedYear,
			Price:         book.Price,
			CreatedAt:     book.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     book.UpdatedAt.Format(time.RFC3339),
		})

	}
	return bookResponses, nil
}
