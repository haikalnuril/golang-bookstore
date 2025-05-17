package routes

import (
	"bookstore/internal/modules/order/controller"
	"bookstore/internal/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewOrderRouter(router fiber.Router, controller *controller.OrderController) {
	order := router.Group("/orders", middleware.Protected())
	order.Post("/", controller.Create)
	order.Get("/", controller.GetAll)
}
