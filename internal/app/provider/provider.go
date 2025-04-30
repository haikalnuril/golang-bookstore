package provider

import (
	"bookstore/internal/app/config"
	"bookstore/internal/modules/user"

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
	// List semua modul
	modules := []interface {
		Register()
		Migrate(db *gorm.DB) error
	}{
		&user.UserModule{App: p.App, DB: p.DB},
		// Tambah modul lain di sini
	}

	// Register routes
	for _, m := range modules {
		m.Register()
	}

	// AutoMigrate entities
	for _, m := range modules {
		if err := m.Migrate(p.DB); err != nil {
			panic("Failed to migrate: " + err.Error())
		}
	}
}
