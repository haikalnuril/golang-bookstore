package postgres

import (
	"bookstore/internal/modules/user/entity"

	"gorm.io/gorm"
)

type userPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *userPostgres {
	return &userPostgres{db: db}
}

func (r *userPostgres) Create(user *entity.User) error {
	return r.db.Create(user).Error // r stand for repository
}

func (r *userPostgres) GetAll() ([]entity.User, error) {
	users := []entity.User{}
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userPostgres) GetByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
