package handler

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/infrastructure/cache"
	"github.com/rbcorrea/meli-test/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

func ResolveShortURL(redis cache.CacheClient, mongo *repository.MongoRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")

		url, err := redis.Get(context.TODO(), code)
		if err != nil {
			url, err = mongo.FindByCode(c.Context(), code)
			if err != nil {
				return c.SendStatus(fiber.StatusNotFound)
			}
			_ = redis.Set(context.TODO(), code, url)
		}
		return c.Redirect(url, fiber.StatusFound)
	}
}
