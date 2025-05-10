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

func (c *BookController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return &exception.BadRequestError{Message: "ID is required"}
	}

	book, err := c.usecase.GetByID(id)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to get book"}
	}

	if book == nil {
		return &exception.NotFoundError{Message: "Book not found"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(book, "Book retrieved successfully"))
}

func (c *BookController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var req model.UpdateBookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
	}

	req.ID = id

	book, err := c.usecase.Update(&req)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to update book"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(book, "Book updated successfully"))

}

func (c *BookController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return &exception.BadRequestError{Message: "ID is required"}
	}

	err := c.usecase.Delete(id)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to delete book"}
	}

	return ctx.Status(fiber.StatusNoContent).JSON(utils.Success(nil, "Book deleted successfully"))
}
