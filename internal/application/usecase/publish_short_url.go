package usecase

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/domain/entity"
)

type Publisher interface {
	PublishShortURL(ctx context.Context, shortURL *entity.ShortURL) error
}

type PublishShortURLUseCase struct {
	Publisher Publisher
}

func NewPublishShortURLUseCase(publisher Publisher) *PublishShortURLUseCase {
	return &PublishShortURLUseCase{Publisher: publisher}
}

func (uc *PublishShortURLUseCase) Execute(ctx context.Context, shortURL *entity.ShortURL) error {
	return uc.Publisher.PublishShortURL(ctx, shortURL)
}
