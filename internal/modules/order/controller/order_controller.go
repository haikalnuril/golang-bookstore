package controller

import (
	"bookstore/internal/modules/order/model"
	"bookstore/internal/modules/order/usecase"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	usecase *usecase.OrderUsecase
}

func NewOrderController(u *usecase.OrderUsecase) *OrderController {
	return &OrderController{usecase: u}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {
	var req model.CreateOrderRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	order, err := c.usecase.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(order)
}

func (c *OrderController) GetAll(ctx *fiber.Ctx) error {
	orders, err := c.usecase.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(orders)
}

func (c *OrderController) Update(ctx *fiber.Ctx) error {
	orderID := ctx.Params("id")
	var req model.UpdateOrderRequest

	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	order, err := c.usecase.Update(orderID, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(order)
}

func (c *OrderController) Delete(ctx *fiber.Ctx) error {
	orderID := ctx.Params("id")

	if err := c.usecase.Delete(orderID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}