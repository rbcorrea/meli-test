package api

import (
	"github.com/rbcorrea/meli-test/internal/application/usecase"
	"github.com/rbcorrea/meli-test/internal/domain/interfaces"
	"github.com/rbcorrea/meli-test/internal/infrastructure/api/handler"

	// "github.com/rbcorrea/meli-test/internal/infrastructure/cache"

	"github.com/rbcorrea/meli-test/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, mongoRepo *repository.MongoRepository, producer interfaces.Producer) {

	shortenUseCase := usecase.NewShortenURLUseCase(producer)
	resolveUseCase := usecase.NewResolveURLUseCase(mongoRepo)
	deleteUseCase := usecase.NewDeleteURLUseCase(mongoRepo)

	app.Post("/shorten", handler.ShortenURL(shortenUseCase))
	app.Get("/:code", handler.ResolveShortURL(resolveUseCase))
	app.Delete("/:code", handler.DeleteURL(deleteUseCase))
}
