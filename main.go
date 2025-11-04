package main

import (
	"go-ecommerce-app/configs"
	"go-ecommerce-app/internal/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENVIROMENT") == "development" {
		godotenv.Load()
	}
	cfg, err := configs.SetUpEnv()

	if err != nil {
		log.Fatalf("configs not loaded: %v\n", err)
	}

	api.StartServer(cfg)

}
