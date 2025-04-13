package usecase

import (
	"context"

	"github.com/rbcorrea/meli-test/internal/application/repository"
	"github.com/rbcorrea/meli-test/internal/domain/interfaces"
)

type DeleteURLUseCase struct {
	Repo repository.URLRepository
}

func NewDeleteURLUseCase(repo repository.URLRepository) interfaces.DeleteURLUseCase {
	return &DeleteURLUseCase{Repo: repo}
}

func (u *DeleteURLUseCase) Execute(ctx context.Context, code string) error {
	return u.Repo.DeactivateByCode(ctx, code)
}
