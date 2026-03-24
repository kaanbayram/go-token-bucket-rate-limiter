package redisclient

import (
	"context"
	"fmt"
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
		panic(err)
	}

	fmt.Println("Redis connected ✅")

	return &RedisClient{
		Rdb: rdb,
	}, nil
}
