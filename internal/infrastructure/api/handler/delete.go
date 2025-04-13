package handler

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/infrastructure/cache"
	"github.com/rbcorrea/meli-test/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

func DeleteURL(redis *cache.Client, mongo *repository.MongoRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")

		_ = redis.Delete(context.TODO(), code)
		_ = mongo.DeactivateByCode(c.Context(), code)

		return c.SendStatus(fiber.StatusNoContent)
	}
}
