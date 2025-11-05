package api

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/routes"

	"github.com/gofiber/fiber/v2"
)

func StartServer(cfg configs.AppConfig) {
	port := cfg.ServerPort
	app := fiber.New()

	v1Routes := app.Group("/v1")

	v1Routes.Get("/health", healthCheck)

	rh := &rest.RestHandler{
		App: v1Routes,
	}

	setupRoutes(rh)

	app.Listen(port)
}

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message": "Server is up and running",
	})
}

func setupRoutes(restHand *rest.RestHandler) {
	routes.UserRoutes(restHand)
}
