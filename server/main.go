package main

import (
	"context"
	"fmt"
	"go-token-bucket-rate-limiter/ratelimiter"
	"go-token-bucket-rate-limiter/redisclient"
	"time"
)

func main() {
	ctx := context.Background()
	redisClient, err := redisclient.New(ctx, "localhost:6379")

	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

	limiter := ratelimiter.New(redisClient, ratelimiter.RateLimiterConfig{
		Capacity:   10,
		RefillRate: 2,
		TTL:        60 * time.Second,
		KeyPrefix:  "rate_limit",
		FailOpen:   true,
	})

}
