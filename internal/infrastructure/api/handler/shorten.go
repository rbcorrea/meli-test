package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-test/internal/domain/interfaces"
)

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

func ShortenURL(shortenURLUseCase interfaces.ShortenURLUseCase) fiber.Handler {
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
			OriginalURL: shortURL.Original,
			ShortURL:    "https://me.li/" + shortURL.Code,
		})
	}
}
