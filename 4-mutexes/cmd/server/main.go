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
	config.LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.GetEnv("DBHost", ""),
		config.GetEnv("DBUser", ""),
		config.GetEnv("DBPassword", ""),
		config.GetEnv("DBName", ""),
		config.GetEnv("DBPort", ""))

	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	clickCounter := counter.NewClickCounter()
	defer clickCounter.Stop()

	app := fiber.New()

	handler.RegisterRoutes(app, database, clickCounter)
}
