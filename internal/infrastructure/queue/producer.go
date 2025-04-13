package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rbcorrea/meli-test/internal/domain/entity"
	"github.com/rbcorrea/meli-test/internal/settings"
	"github.com/rs/zerolog/log"
)

type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewProducer(cfg *settings.Config) (*Producer, error) {
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to Queue")
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Error().Err(err).Msg("Error opening channel")
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		cfg.RabbitMQQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Err(err).Msg("Error declaring queue")
		return nil, err
	}

	log.Info().Str("queue", queue.Name).Msg("Queue declared")

	return &Producer{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}, nil
}

func (p *Producer) PublishShortenURL(ctx context.Context, message *entity.ShortURL) error {
	body, err := json.Marshal(message)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling message")
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = p.channel.PublishWithContext(ctx,
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Error().Err(err).Msg("Error publishing message")
		return fmt.Errorf("Error publishing message: %w", err)
	}

	log.Info().
		Str("queue", p.queue.Name).
		Bytes("body", body).
		Msg("Message published")

	return nil
}

func (p *Producer) Close() {
	if p.channel != nil {
		_ = p.channel.Close()
	}
	if p.conn != nil {
		_ = p.conn.Close()
	}
}
