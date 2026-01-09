package api

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/routes"
	"go-ecommerce-app/internal/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(cfg configs.AppConfig) {
	port := cfg.ServerPort
	app := fiber.New()

	auth := helper.InitializeAuth(cfg.Secret)

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})

	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	log.Printf("Database Connected")
	helper.RunExecs(db)
	err = helper.RunMigrations(db)

	if err != nil {
		log.Fatalf("database migrations error: %v", err)
	}

	c := cors.New(cors.Config{
		AllowOrigins: "http://localhost:3030",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	})

	app.Use(c)

	v1Routes := app.Group("/api/v1")

	v1Routes.Get("/health", healthCheck)

	rh := &rest.RestHandler{
		App:           v1Routes,
		DB:            db,
		Auth:          auth,
		Configuration: cfg,
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
	routes.CatalogRoutes(restHand)
}