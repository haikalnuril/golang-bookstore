package user

import (
	"bookstore/internal/modules/user/controller"
	"bookstore/internal/modules/user/entity"
	"bookstore/internal/modules/user/repository/postgres"
	"bookstore/internal/modules/user/router"
	"bookstore/internal/modules/user/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserModule struct {
	App *fiber.App
	DB  *gorm.DB
}

func (u *UserModule) Register() {
	userRepo := postgres.NewUserPostgres(u.DB)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	api := u.App.Group("/api/v1")
	router.NewUserRouter(api, userController)
}

func (u *UserModule) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}
