package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-test/internal/domain/interfaces"
)

func ResolveShortURL(resolveURLUseCase interfaces.ResolveURLUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")
		if code == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "code is required",
			})
		}

		shortURL, err := resolveURLUseCase.Execute(c.Context(), code)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "URL not found",
			})
		}

		return c.Redirect(shortURL.Original, fiber.StatusFound)
	}
}
