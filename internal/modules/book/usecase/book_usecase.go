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

func (b *BookUsecase) GetByID(id string) (*model.BookResponse, error) {
	book, err := b.repo.GetByID(id)
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

func (b *BookUsecase) Update(req *model.UpdateBookRequest) (*model.BookResponse, error) {
	updates := make(map[string]interface{})

	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Author != nil {
		updates["author"] = *req.Author
	}
	if req.Genre != nil {
		updates["genre"] = *req.Genre
	}
	if req.PublishedYear != nil {
		updates["published_year"] = *req.PublishedYear
	}
	if req.Price != nil {
		updates["price"] = *req.Price
	}

	updatedBook, err := b.repo.Update(req.ID, updates)
	if err != nil {
		return nil, err
	}

	return &model.BookResponse{
		ID:            updatedBook.ID.String(),
		Title:         updatedBook.Title,
		Author:        updatedBook.Author,
		Genre:         updatedBook.Genre,
		PublishedYear: updatedBook.PublishedYear,
		Price:         updatedBook.Price,
		CreatedAt:     updatedBook.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     updatedBook.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (b *BookUsecase) Delete(id string) error {
	return b.repo.Delete(id)
}
