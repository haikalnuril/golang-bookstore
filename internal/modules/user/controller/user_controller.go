package controller

import (
	"bookstore/internal/app/config"
	"bookstore/internal/app/exception"
	"bookstore/internal/app/utils"
	"bookstore/internal/modules/user/model"
	"bookstore/internal/modules/user/usecase"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase   *usecase.UserUsecase
	validator *config.Validator
}

func NewUserController(usecase *usecase.UserUsecase, validator *config.Validator) *UserController {
	return &UserController{
		usecase:   usecase,
		validator: validator,
	}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	var req model.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
	}

	if errs := c.validator.Validate(req); len(errs) > 0 {
		return &exception.BadRequestError{Message: c.validator.Message(errs)}
	}

	user, err := c.usecase.Create(&req)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to create user"}
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.Created(user, "User created successfully"))
}

func (c *UserController) GetAll(ctx *fiber.Ctx) error {
	users, err := c.usecase.GetAll()
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to get users"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(users, "Users retrieved successfully"))
}

func (c *UserController) GetByEmail(ctx *fiber.Ctx) error {
	email := ctx.Params("email")
	if email == "" {
		return &exception.BadRequestError{Message: "Email is required"}
	}

	user, err := c.usecase.GetByEmail(email)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to get user"}
	}

	if user == nil {
		return &exception.NotFoundError{Message: "User not found"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(user, "User retrieved successfully"))
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req model.UserUpdateRequest

	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
	}

	req.ID = id

	if errs := c.validator.Validate(req); len(errs) > 0 {
		return &exception.BadRequestError{Message: c.validator.Message(errs)}
	}

	user, err := c.usecase.Update(&req)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to update user"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(user, "User updated successfully"))
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return &exception.BadRequestError{Message: "ID is required"}
	}

	err := c.usecase.Delete(id)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to delete user"}
	}

	return ctx.Status(fiber.StatusNoContent).JSON(utils.Success(nil, "User deleted successfully"))
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var req model.UserLoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
	}

	if errs := c.validator.Validate(req); len(errs) > 0 {
		return &exception.BadRequestError{Message: c.validator.Message(errs)}
	}

	user, err := c.usecase.Login(&req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "invalid password") {
			return &exception.UnauthorizedError{Message: "Email or password is incorrect"}
		}

		fmt.Printf("Login error: %v\n", err)
		return &exception.InternalServerError{Message: "Authentication failed"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(user, "Login successful"))
}
