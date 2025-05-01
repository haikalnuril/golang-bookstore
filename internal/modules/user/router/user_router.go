package router

import (
	"bookstore/internal/modules/user/controller"

	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(router fiber.Router, controller *controller.UserController) {
	user := router.Group("/users")
	user.Post("/", controller.Create) // POST /users
	user.Get("/", controller.GetAll) // GET /users
	user.Get("/:email", controller.GetByEmail) // GET /users/:email
	user.Put("/:id", controller.Update) // PUT /users/:id
	user.Delete("/:id", controller.Delete) // DELETE /users/:id
}
