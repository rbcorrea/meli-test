package interfaces

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/domain/entity"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type Repository interface {
	FindByCode(ctx context.Context, code string) (string, error)
}

type ShortenURLUseCase interface {
	Execute(ctx context.Context, originalURL string) (*entity.ShortURL, error)
}

type ResolveURLUseCase interface {
	Execute(ctx context.Context, code string) (*entity.ShortURL, error)
}

type DeleteURLUseCase interface {
	Execute(ctx context.Context, code string) error
}

type Producer interface {
	PublishShortenURL(ctx context.Context, message *entity.ShortURL) error
}
