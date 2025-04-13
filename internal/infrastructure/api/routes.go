package api

import (
	"log"

	"github.com/rbcorrea/meli-test/internal/application/usecase"
	"github.com/rbcorrea/meli-test/internal/infrastructure/api/handler"
	"github.com/rbcorrea/meli-test/internal/infrastructure/cache"
	"github.com/rbcorrea/meli-test/internal/infrastructure/queue"
	"github.com/rbcorrea/meli-test/internal/infrastructure/repository"
	"github.com/rbcorrea/meli-test/internal/settings"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	config := settings.Load()

	redisClient := cache.NewRedisClient(config.RedisAddr)
	producer, err := queue.NewProducer(config)
	if err != nil {
		log.Fatal("Producer Error: %v", err)
	}
	shortenUseCase := usecase.NewShortenURLUseCase(producer)

	app.Get("/health", handler.HealthCheck)

	app.Post("/shorten", handler.ShortenURL(producer, shortenUseCase))
	app.Get("/:code", handler.ResolveShortURL(redisClient, repository.NewMongoRepository()))
	// app.Get("/stats/:code", handler.GetStats(mongoRepo))
	app.Delete("/:code", handler.DeleteURL(redisClient, mongoRepo))
}
