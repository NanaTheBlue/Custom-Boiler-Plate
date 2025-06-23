package redis

import (
	"context"
	"time"
)

type Store interface {
	Add(ctx context.Context, key string, value []byte, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Subscribe(ctx context.Context, channel string, handler func(message string)) error
	Publish(ctx context.Context, channel string, message []byte) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Count(ctx context.Context, key string) (int64, error)
}
