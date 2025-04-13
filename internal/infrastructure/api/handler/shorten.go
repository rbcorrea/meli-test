package handler

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/infrastructure/queue"

	"github.com/gofiber/fiber/v2"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

type ShortenURLUseCase interface {
	Execute(ctx context.Context, url string) (ShortenedURL, error)
}

type ShortenedURL struct {
	OriginalURL string
	Code        string
}

func ShortenURL(producer *queue.Producer, shortenURLUseCase ShortenURLUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req shortenRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		if req.URL == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "URL is required",
			})
		}

		shortURL, err := shortenURLUseCase.Execute(c.Context(), req.URL)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to shorten URL",
			})
		}

		return c.Status(fiber.StatusOK).JSON(shortenResponse{
			OriginalURL: shortURL.OriginalURL,
			ShortURL:    "https://me.li/" + shortURL.Code,
		})
	}
}
