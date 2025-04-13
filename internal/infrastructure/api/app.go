package api

import (
	"github.com/rbcorrea/meli-test/internal/domain/interfaces"
	"github.com/rbcorrea/meli-test/internal/infrastructure/api/middleware"
	"github.com/rbcorrea/meli-test/internal/infrastructure/repository"

	// "github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
)

func NewApp(mongoRepo *repository.MongoRepository, producer interfaces.Producer) *fiber.App {
	app := fiber.New()

	// app.Use(otelfiber.Middleware())
	app.Use(middleware.Logger())

	RegisterRoutes(app, mongoRepo, producer)

	return app
}
