package order

import (
	bookpostgres "bookstore/internal/modules/book/repository/postgres"
	"bookstore/internal/modules/order/controller"
	"bookstore/internal/modules/order/entity"
	"bookstore/internal/modules/order/repository/postgres"
	"bookstore/internal/modules/order/router"
	"bookstore/internal/modules/order/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OrderModule struct {
	App *fiber.App
	DB  *gorm.DB
}

func (o *OrderModule) Register() {
	orderRepo := postgres.NewOrderPostgres(o.DB)
	bookRepo := bookpostgres.NewBookPostgres(o.DB)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, bookRepo)
	orderController := controller.NewOrderController(orderUsecase)

	api := o.App.Group("/api/v1")
	router.NewOrderRouter(api, orderController)
}

func (o *OrderModule) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Order{})
}
