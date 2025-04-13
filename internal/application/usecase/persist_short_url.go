package usecase

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/application/repository"
	"github.com/rbcorrea/meli-test/internal/domain/entity"
)

type PersistShortURLUseCase struct {
	Repository repository.URLRepository
}

func NewPersistShortURLUseCase(repo repository.URLRepository) *PersistShortURLUseCase {
	return &PersistShortURLUseCase{Repository: repo}
}

func (uc *PersistShortURLUseCase) Execute(ctx context.Context, shortURL *entity.ShortURL) error {
	return uc.Repository.Save(ctx, shortURL)
}
