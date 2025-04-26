package main

import (
	"bookstore/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()

	config.ConnPostgres(cfg)

	app.Listen(":" + cfg.Port)
}
