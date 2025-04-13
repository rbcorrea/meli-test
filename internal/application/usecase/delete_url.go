package usecase

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/application/repository"
)

type DeleteURLUseCase struct {
	Repo repository.URLRepository
}

func NewDeleteURLUseCase(repo repository.URLRepository) *DeleteURLUseCase {
	return &DeleteURLUseCase{Repo: repo}
}

func (u *DeleteURLUseCase) Execute(ctx context.Context, code string) error {
	return u.Repo.DeactivateByCode(ctx, code)
}
