package api

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/routes"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/schema"
	"log"

	"github.com/gofiber/fiber/v2"
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

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec(`
			DO $$
			BEGIN
				IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_type') THEN
					CREATE TYPE payment_type AS ENUM ('daily', 'weekly', 'monthly');
				END IF;
			END$$;
			`)

	err = db.AutoMigrate(&schema.User{})
	if err != nil {
		log.Fatalf("error running migrations: %v", err)
	}

	err = db.AutoMigrate(&schema.BankAccount{})
	if err != nil {
		log.Fatalf("error running migrations: %v", err)
	}

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
}
