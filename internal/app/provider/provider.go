package provider

import (
	"bookstore/internal/app/config"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Provider struct {
	App       *fiber.App
	Config    *config.Config
	DB        *gorm.DB
	Validator *config.Validator
}

func (p *Provider) Provide() {
}
