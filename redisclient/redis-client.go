package redisclient

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Rdb *redis.Client
}

func New(ctx context.Context, url string) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: url,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{
		Rdb: rdb,
	}, nil
}
