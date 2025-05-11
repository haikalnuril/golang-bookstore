package postgres

import (
	"bookstore/internal/modules/book/entity"
	"bookstore/internal/modules/book/model"

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

func (r *bookPostgres) GetAll(req *model.SearchBookRequest) ([]entity.Book, int, error) {
	var books []entity.Book
	query := r.db.Model(&entity.Book{})

	if req.Title != nil {
		query = query.Where("title ILIKE ?", "%"+*req.Title+"%")
	}
	if req.Author != nil {
		query = query.Where("author ILIKE ?", "%"+*req.Author+"%")
	}
	if req.Genre != nil {
		query = query.Where("genre ILIKE ?", "%"+*req.Genre+"%")
	}
	if req.PublishedYear != nil {
		query = query.Where("published_year = ?", *req.PublishedYear)
	}
	if req.Price != nil {
		query = query.Where("price = ?", *req.Price)
	}

	var total int64
	query.Count(&total)

	offset := (req.Page - 1) * req.Size
	err := query.Limit(req.Size).Offset(offset).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	return books, int(total), nil
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
