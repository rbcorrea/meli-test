package repository

import (
	"context"
	"time"

	"github.com/rbcorrea/meli-test/internal/domain/entity"
)

type URLRepository interface { //mudar para URL apenas
	Save(ctx context.Context, shortURL *entity.ShortURL) error
	FindByCode(ctx context.Context, code string) (*entity.ShortURL, error)
	UpdateAccessData(ctx context.Context, code string, accessed int64, lastAccess time.Time) error
	DeactivateByCode(ctx context.Context, code string) error
}
