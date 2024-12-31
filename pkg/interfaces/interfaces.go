package interfaces

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

// DB is an interface for database operations.
type DB interface {
	PingContext(ctx context.Context) error
	Close() error
}

// Redis is an interface for Redis operations.
type Redis interface {
	Ping(ctx context.Context) *redis.StatusCmd
	Close() error
}

// KafkaReader is an interface for Kafka reader operations.
type KafkaWriter interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}
