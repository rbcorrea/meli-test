package main

import (
	"context"
	"log"

	"github.com/rbcorrea/meli-test/internal/infrastructure/api"
	"github.com/rbcorrea/meli-test/internal/infrastructure/cache"

	// "github.com/rbcorrea/meli-test/internal/infrastructure/observability"
	"github.com/rbcorrea/meli-test/internal/infrastructure/queue"
	"github.com/rbcorrea/meli-test/internal/infrastructure/repository"
	"github.com/rbcorrea/meli-test/internal/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := settings.Load()

	redis := cache.NewRedisClient(cfg.RedisAddr)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	mongoDatabase := client.Database("url_shortener")
	mongoRepo := repository.NewMongoRepository(mongoDatabase)

	_ = queue.StartConsumer(cfg.RabbitMQURL, mongoRepo, redis)

	app := api.NewApp()

	log.Fatal(app.Listen(":8080"))
}
