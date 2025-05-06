package router

import (
	"bookstore/internal/app/middleware"
	"bookstore/internal/modules/user/controller"

	"github.com/gofiber/fiber/v2"
)

func NewUserRouter(router fiber.Router, controller *controller.UserController) {
	// Public routes
	auth := router.Group("/auth")
	auth.Post("/login", controller.Login)     // POST /auth/login
	auth.Post("/register", controller.Create) // POST /auth/register

	// Protected routes
	user := router.Group("/users", middleware.Protected())
	user.Get("/", controller.GetAll)           // GET /users
	user.Get("/:email", controller.GetByEmail) // GET /users/:email
	user.Put("/:id", controller.Update)        // PUT /users/:id
	user.Delete("/:id", controller.Delete)     // DELETE /users/:id
}
