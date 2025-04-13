package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-test/internal/domain/interfaces"
)

func DeleteURL(deleteURLUseCase interfaces.DeleteURLUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")
		if code == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "code is required",
			})
		}

		err := deleteURLUseCase.Execute(c.Context(), code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to delete URL",
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
