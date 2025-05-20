package handler

import (
	"mutexExercise/counter"
	"mutexExercise/internal/db"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, database *gorm.DB, clickCounter *counter.ClickCounter) {
	app.Post("/shorten", func(c *fiber.Ctx) error {
		type Request struct {
			URL string `json:"url"`
		}

		var body Request
		if err := c.BodyParser(&body); err != nil {
			return fiber.ErrBadRequest
		}

		shortURL, err := db.CreateShortURL(database, body.URL)
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(fiber.Map{
			"short_code": shortURL.ShortCode,
		})
	})

	app.Get("/:code", func(c *fiber.Ctx) error {
		code := c.Params("code")
		shortURL, err := db.GetShortURL(database, code)
		if err != nil {
			return fiber.ErrNotFound
		}

		// Increment in-memory click counter
		clickCounter.Increment(code)

		// Increment click count in DB asynchronously (optional)
		go database.Model(&shortURL).UpdateColumn("click_count", gorm.Expr("click_count + ?", 1))

		return c.Redirect(shortURL.OriginalURL, http.StatusFound)
	})
}
