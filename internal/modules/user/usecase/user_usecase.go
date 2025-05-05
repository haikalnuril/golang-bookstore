package usecase

import (
	"bookstore/internal/app/utils"
	"bookstore/internal/modules/user/entity"
	"bookstore/internal/modules/user/model"
	"bookstore/internal/modules/user/repository"
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Create(req *model.UserRequest) (*model.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	err = u.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return &model.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserUsecase) GetAll() ([]model.UserResponse, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var userResponses []model.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, model.UserResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return userResponses, nil
}

func (u *UserUsecase) GetByEmail(email string) (*model.UserResponse, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return &model.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserUsecase) Update(req *model.UserUpdateRequest) (*model.UserResponse, error) {
	user := &entity.User{
		Name:  req.Name,
		Email: req.Email,
	}

	updatedUser, err := u.repo.Update(req.ID, user)
	if err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:    updatedUser.ID.String(),
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	}, nil
}

func (u *UserUsecase) Delete(id string) error {
	return u.repo.Delete(id)

}

func (u *UserUsecase) Login(req *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	user, err := u.repo.GetByEmail(req.Email)
	
	// Check if it's a "not found" error from the database
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found - return a more specific error
			return nil, fmt.Errorf("user with email %s not found", req.Email)
		}
		// Some other database error occurred
		return nil, fmt.Errorf("database error: %w", err)
	}
	
	// User found, now check password
	if !utils.ComparePassword(user.Password, req.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	// Generate JWT token
	return &model.UserLoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: user.ID.String(), // Consider using a proper JWT token here
	}, nil
}
