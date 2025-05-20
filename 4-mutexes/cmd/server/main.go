package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"mutexExercise/counter"
	"mutexExercise/internal/config"
	"mutexExercise/internal/db"
	"mutexExercise/internal/handler"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	clickCounter := counter.NewClickCounter(cfg)
	defer clickCounter.Stop()

	app := fiber.New()

	handler.RegisterRoutes(app, database, clickCounter)

	fmt.Printf("Server running at http://localhost:%d\n", cfg.ServerPort)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.ServerPort)))
}
