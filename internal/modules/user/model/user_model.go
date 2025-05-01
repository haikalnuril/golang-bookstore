package model

type UserRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserPageResponse struct {
	Page  int
	Size  int
	Total int
	User  []UserResponse
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserUpdateRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"min=2"`
	Email string `json:"email" validate:"email"`
}
