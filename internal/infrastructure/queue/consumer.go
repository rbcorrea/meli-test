package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
	"github.com/rbcorrea/meli-test/internal/domain/entity"
	"github.com/rbcorrea/meli-test/internal/infrastructure/repository"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ShortURLMessage struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

func StartConsumer(url string, repo *repository.MongoRepository, redis *redis.Client) error {

	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare("shorten_url_queue", true, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {

			var payload ShortURLMessage
			shortURLDocument := entity.NewShortURL(payload.URL, payload.Code)

			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}
			_ = redis.Set(context.TODO(), payload.Code, payload.URL)
			_ = repo.Save(context.TODO(), shortURLDocument)
		}
	}()
	return nil
}
