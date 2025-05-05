package handlers

import (
	"exercise3/config"
	"exercise3/crawler"
	"exercise3/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"sync"
)

func CrawlHandler(c *fiber.Ctx) error {
	type Requests struct {
		URLs []string `json:"urls"`
	}

	var req Requests
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body.")
	}

	var wg sync.WaitGroup
	for _, url := range req.URLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			results := crawler.Crawl([]string{url})
			for _, r := range results {
				page := models.CrawlPage{
					URL:      r.URL,
					Elements: r.Elements,
				}
				if err := config.DB.Create(&page).Error; err != nil {
					fmt.Println("DB error:", err)
				}
			}
		}(url)
	}
	wg.Wait()

	return c.JSON(fiber.Map{
		"message": "Crawl completed.",
	})
}
