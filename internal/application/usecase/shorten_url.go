package usecase

import (
	"context"

	"math/rand"
	"time"

	"github.com/rbcorrea/meli-test/internal/domain/entity"
	"github.com/rbcorrea/meli-test/internal/infrastructure/queue"
)

type ShortenURLUseCase struct {
	Producer queue.Producer
}

func NewShortenURLUseCase(producer queue.Producer) *ShortenURLUseCase {
	return &ShortenURLUseCase{
		Producer: producer,
	}
}
func GenerateCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, n)
	for i := range code {
		code[i] = letters[r.Intn(len(letters))]
	}
	return string(code)
}

func (u *ShortenURLUseCase) Execute(ctx context.Context, originalURL string) (*entity.ShortURL, error) {
	code := GenerateCode(6)
	shortURL := entity.NewShortURL(originalURL, code)

	err := u.Producer.PublishShortenURL(ctx, shortURL)
	if err != nil {
		return nil, err
	}

	return shortURL, nil
}
