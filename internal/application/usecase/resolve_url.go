package usecase

import (
	"context"
	"errors"

	"github.com/rbcorrea/meli-test/internal/application/repository"
	"github.com/rbcorrea/meli-test/internal/domain/entity"
)

type ResolveURLUseCase struct {
	Repo repository.URLRepository
}

func NewResolveURLUseCase(repo repository.URLRepository) *ResolveURLUseCase {
	return &ResolveURLUseCase{Repo: repo}
}

func (u *ResolveURLUseCase) Execute(ctx context.Context, code string) (*entity.ShortURL, error) {
	shortURL, err := u.Repo.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	if shortURL == nil || !shortURL.IsActive {
		return nil, errors.New("short URL not found or inactive")
	}

	// shortURL.IncrementAccess()
	// _ = u.Repo.UpdateAccessData(ctx, code, shortURL.Accessed, shortURL.LastAccess)

	return shortURL, nil
}
