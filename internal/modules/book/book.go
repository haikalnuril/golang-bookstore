package book

import (
	"bookstore/internal/modules/book/controller"
	"bookstore/internal/modules/book/entity"
	"bookstore/internal/modules/book/repository/postgres"
	"bookstore/internal/modules/book/router"
	"bookstore/internal/modules/book/usecase"
	"bookstore/internal/app/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookModule struct {
	App *fiber.App
	DB  *gorm.DB
}

func (b *BookModule) Register() {
	bookRepo := postgres.NewBookPostgres(b.DB)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	validator := config.NewValidator()
	bookController := controller.NewBookController(bookUsecase, validator)

	api := b.App.Group("/api/v1")
	router.NewBookRouter(api, bookController)
}

func (b *BookModule) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Book{})
}
