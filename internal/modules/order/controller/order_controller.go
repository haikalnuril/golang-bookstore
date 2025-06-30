package controller

import (
	"bookstore/internal/app/config"
	"bookstore/internal/app/exception"
	"bookstore/internal/app/utils"
	"bookstore/internal/modules/order/model"
	"bookstore/internal/modules/order/usecase"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	usecase   *usecase.OrderUsecase
	validator *config.Validator
}

func NewOrderController(u *usecase.OrderUsecase, v *config.Validator) *OrderController {
	return &OrderController{
		usecase:   u,
		validator: v,
	}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {
	var req model.CreateOrderRequest
	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request"}
	}

	if errs := c.validator.Validate(req); len(errs) > 0 {
		return &exception.BadRequestError{Message: c.validator.Message(errs)}
	}

	order, err := c.usecase.Create(req)
	if err != nil {
		return &exception.InternalServerError{Message: err.Error()}
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.Created(order, "Order created successfully"))
}

func (c *OrderController) GetAll(ctx *fiber.Ctx) error {
	orders, err := c.usecase.GetAll()
	if err != nil {
		return &exception.InternalServerError{Message: err.Error()}
	}
	return ctx.Status(fiber.StatusOK).JSON(utils.Success(orders, "Orders retrieved successfully"))
}

func (c *OrderController) Update(ctx *fiber.Ctx) error {
	orderID := ctx.Params("id")
	var req model.UpdateOrderRequest

	if err := ctx.BodyParser(&req); err != nil {
		return &exception.BadRequestError{Message: "Invalid request"}
	}

	if errs := c.validator.Validate(req); len(errs) > 0 {
		return &exception.BadRequestError{Message: c.validator.Message(errs)}
	}

	order, err := c.usecase.Update(orderID, req)
	if err != nil {
		return &exception.InternalServerError{Message: err.Error()}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.Success(order, "Order updated successfully"))
}

func (c *OrderController) Delete(ctx *fiber.Ctx) error {
	orderID := ctx.Params("id")

	if err := c.usecase.Delete(orderID); err != nil {
		return &exception.InternalServerError{Message: err.Error()}
	}

	return ctx.Status(fiber.StatusNoContent).JSON(utils.Success(nil, "Order deleted successfully"))
}
