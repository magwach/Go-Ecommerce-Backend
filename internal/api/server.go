package api

import (
	"go-ecommerce-app/configs"

	"github.com/gofiber/fiber/v2"
)

func StartServer(cfg configs.AppConfig) {
	port := cfg.ServerPort
	app := fiber.New()
	app.Listen(port)
}
