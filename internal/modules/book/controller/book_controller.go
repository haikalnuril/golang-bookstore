package controller

import (
	"bookstore/internal/app/config"
	"bookstore/internal/app/exception"
	"bookstore/internal/app/utils"
	"bookstore/internal/modules/book/model"
	"bookstore/internal/modules/book/usecase"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	usecase   *usecase.BookUsecase
	validator *config.Validator
}

func NewBookController(usecase *usecase.BookUsecase, validator *config.Validator) *BookController {
	return &BookController{
		usecase:   usecase,
		validator: validator,
	}
}

func (c *BookController) Create(ctx *fiber.Ctx) error {
	var req model.BookRequest
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request body"}
	}

	if errs := c.validator.Validate(req); len(errs) > 0 {
		return &exception.BadRequestError{Message: c.validator.Message(errs)}
	}

	book, err := c.usecase.Create(&req)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to create book"}
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.Created(book, "Book created successfully"))
}

func (c *BookController) GetAll(ctx *fiber.Ctx) error {
	var req model.SearchBookRequest

	if err := ctx.QueryParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid query parameters"}
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 10
	}

	pageable, err := c.usecase.GetAll(&req)
	if err != nil {
		return &exception.InternalServerError{Message: "Failed to get books"}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(pageable, "Books retrieved successfully"))
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

	if errs := c.validator.Validate(req); len(errs) > 0 {
		return &exception.BadRequestError{Message: c.validator.Message(errs)}
	}

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
