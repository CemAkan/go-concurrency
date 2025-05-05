package main

import (
	"exercise3/config"
	"exercise3/handlers"
	"exercise3/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	utils.LoadEnv()
	config.ConnectDatabase()

	app := fiber.New()

	app.Post("/crawl", handlers.CrawlHandler)

	port := utils.GetEnv("APP_PORT", "1234")

	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Port can not listening")
	}
}
