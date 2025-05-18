package usecase

import (
	"bookstore/internal/modules/order/entity"
	"bookstore/internal/modules/order/model"
	"bookstore/internal/modules/order/repository"
	bookrepo "bookstore/internal/modules/book/repository"
	"time"

	"github.com/google/uuid"
)

type OrderUsecase struct {
	repo repository.OrderRepository
	bookRepo        bookrepo.BookRepository
}

func NewOrderUsecase(repo repository.OrderRepository, bookRepo bookrepo.BookRepository) *OrderUsecase {
	return &OrderUsecase{
		repo:    repo,
		bookRepo: bookRepo,
	}
}

func (u *OrderUsecase) Create(req model.CreateOrderRequest) (*model.OrderResponse, error) {
	order := &entity.Order{
		ID:        uuid.New(),
		UserID:    req.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	total := 0
	for _, item := range req.Items {
		book, err := u.bookRepo.GetByID(item.BookID.String())
		if err != nil {
			return nil, err
		}

		unitPrice := book.Price
		total += unitPrice * item.Quantity

		orderItem := entity.OrderItem{
			ID:        uuid.New(),
			OrderID:   order.ID,
			BookID:    book.ID,
			Quantity:  item.Quantity,
			UnitPrice: unitPrice,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		order.Items = append(order.Items, orderItem)
	}

	order.TotalPrice = total

	if err := u.repo.Create(order); err != nil {
		return nil, err
	}

	return u.toResponse(order), nil
}

func (u *OrderUsecase) GetAll() ([]model.OrderResponse, error) {
	orders, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []model.OrderResponse
	for _, order := range orders {
		responses = append(responses, *u.toResponse(&order))
	}
	return responses, nil
}

func (u *OrderUsecase) toResponse(order *entity.Order) *model.OrderResponse {
	var items []model.OrderItemResponse
	for _, item := range order.Items {
		items = append(items, model.OrderItemResponse{
			BookID:    item.BookID.String(),
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}
	return &model.OrderResponse{
		ID:         order.ID.String(),
		UserID:     order.UserID.String(),
		TotalPrice: order.TotalPrice,
		Items:      items,
		CreatedAt:  order.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  order.UpdatedAt.Format(time.RFC3339),
	}
}

func (u *OrderUsecase) Update(orderID string, req model.UpdateOrderRequest) (*model.OrderResponse, error) {
	existing, err := u.repo.GetByID(orderID)
	if err != nil {
		return nil, err
	}

	updatedItems := []entity.OrderItem{}
	total := 0

	for _, item := range req.Items {
		book, err := u.bookRepo.GetByID(item.BookID)
		if err != nil {
			return nil, err
		}

		unitPrice := book.Price
		total += unitPrice * item.Quantity

		updatedItems = append(updatedItems, entity.OrderItem{
			ID:        uuid.New(),
			OrderID:   existing.ID,
			BookID:    book.ID,
			Quantity:  item.Quantity,
			UnitPrice: unitPrice,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	existing.Items = updatedItems
	existing.TotalPrice = total
	existing.UpdatedAt = time.Now()

	if err := u.repo.Update(existing); err != nil {
		return nil, err
	}

	return u.toResponse(existing), nil
}

func (u *OrderUsecase) Delete(orderID string) error {
	return u.repo.Delete(orderID)
}
