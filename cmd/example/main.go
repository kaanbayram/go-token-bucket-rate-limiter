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

	limiter, err := ratelimiter.New(redisClient.Rdb, ratelimiter.RateLimiterConfig{
		Capacity:   10,
		RefillRate: 2,
		TTL:        60 * time.Second,
		KeyPrefix:  "rate_limit",
		FailOpen:   true,
	})
	if err != nil {
		fmt.Println("Error creating limiter:", err)
		return
	}

	allowed, err := limiter.Allow("127.0.0.1")
	if err != nil {
		fmt.Println("Allow check failed:", err)
		return
	}
	fmt.Println("Allowed:", allowed)
}
