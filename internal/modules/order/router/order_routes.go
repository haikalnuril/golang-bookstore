package router

import (
	"bookstore/internal/modules/order/controller"
	"bookstore/internal/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewOrderRouter(router fiber.Router, controller *controller.OrderController) {
	order := router.Group("/orders", middleware.Protected())
	order.Post("/", controller.Create)
	order.Get("/", controller.GetAll)
	order.Put("/:id", controller.Update)
	order.Delete("/:id", controller.Delete)
}
