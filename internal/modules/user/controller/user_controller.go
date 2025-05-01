package controller

import (
	"bookstore/internal/app/exception"
	"bookstore/internal/app/utils"
	"bookstore/internal/modules/user/model"
	"bookstore/internal/modules/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	var req model.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
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
	id := ctx.Params("id") // ambil ID dari URL

	var req model.UserUpdateRequest
	
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
	}

	req.ID = id // set ID dari URL ke request body

	user, err := c.usecase.Update(&req)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to update user"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(user, "User updated successfully"))

}