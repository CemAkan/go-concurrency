package main

import (
	"fmt"
	"log"
	"mutexExercise/internal/config"
	"mutexExercise/internal/db"
	"mutexExercise/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	app := fiber.New()

	handler.RegisterRoutes(app, database)

	fmt.Printf("Server running at http://localhost:%d\n", cfg.ServerPort)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.ServerPort)))
}
