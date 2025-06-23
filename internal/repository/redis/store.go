package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type store struct {
	client *redis.Client
}

func NewRedisInstance(redis *redis.Client) Store {
	return &store{redis}
}

func (s *store) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := s.client.Expire(ctx, key, expiration).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *store) Count(ctx context.Context, key string) (int64, error) {
	count, err := s.client.SCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *store) Delete(ctx context.Context, key string) error {
	err := s.client.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *store) Add(ctx context.Context, key string, value []byte, expiration time.Duration) error {

	err := s.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)

	}
	return nil
}

func (s *store) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		panic(err)

	}
	fmt.Println("foo", val)
	return val, nil
}

func (s *store) Publish(ctx context.Context, channel string, message []byte) error {

	err := s.client.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *store) Subscribe(ctx context.Context, channel string, handler func(message string)) error {
	pubsub := s.client.Subscribe(ctx, channel)

	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		handler(msg.Payload)
		if err := ctx.Err(); err != nil {
			return err
		}
	}

	return nil
}
