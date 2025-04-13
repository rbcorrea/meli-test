package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logger.Info().
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("status", c.Response().StatusCode()).
			Dur("duration", duration).
			Msg("incoming request")

		return err
	}
}
