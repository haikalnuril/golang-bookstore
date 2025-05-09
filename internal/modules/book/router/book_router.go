package router

import (
	"bookstore/internal/app/middleware"
	"bookstore/internal/modules/book/controller"

	"github.com/gofiber/fiber/v2"
)

func NewBookRouter(router fiber.Router, controller *controller.BookController) {

	book := router.Group("/books", middleware.Protected())
	book.Get("/", controller.GetAll)  // GET /books
	book.Post("/", controller.Create) // POST /books
	book.Get("/:id", controller.GetByID) // GET /books/:id
	book.Put("/:id", controller.Update) // PUT /books/:id
}
