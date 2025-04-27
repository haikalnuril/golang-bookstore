package main

import (
	"bookstore/internal/app/config"
	"bookstore/internal/app/provider"
	"fmt"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		cfg       = config.LoadConfig()
		fiber     = config.NewFiber(cfg)
		postgres  = config.ConnPostgres(cfg)
		validator = config.NewValidator()
		bootstrap = provider.Provider{App: fiber, Config: cfg, DB: postgres, Validator: validator}
	)

	bootstrap.Provide()

	port, atoiErr := strconv.Atoi(cfg.Port)
	if atoiErr != nil {
		fmt.Println("Error converting port:", atoiErr)
		return
	}
	if err := fiber.Listen(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
