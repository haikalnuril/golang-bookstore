package controller

import (
	"bookstore/internal/app/exception"
	"bookstore/internal/app/utils"
	"bookstore/internal/modules/book/model"
	"bookstore/internal/modules/book/usecase"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	usecase *usecase.BookUsecase
}

func NewBookController(usecase *usecase.BookUsecase) *BookController {
	return &BookController{usecase: usecase}
}

func (c *BookController) Create(ctx *fiber.Ctx) error {
	var req model.BookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
	}

	book, err := c.usecase.Create(&req)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to create book"}
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.Created(book, "Book created successfully"))
}

func (c *BookController) GetAll(ctx *fiber.Ctx) error {
	books, err := c.usecase.GetAll()
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to get books"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(books, "Books retrieved successfully"))
}
