package repository

import "bookstore/internal/modules/user/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetAll() ([]entity.User, error)
	GetByEmail(email string) (*entity.User, error)
}