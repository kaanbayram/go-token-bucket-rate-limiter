package ratelimiter

import (
	"context"
	"fmt"
	"go-token-bucket-rate-limiter/redisclient"
	"go-token-bucket-rate-limiter/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiterConfig struct {
	Capacity   int64
	RefillRate int64
	TTL        time.Duration
	KeyPrefix  string
	FailOpen   bool
}

type RateLimiter struct {
	redisClient    redisclient.RedisClient
	redisLuaScript *redis.Script
	cfg            RateLimiterConfig
}

func New(redisClient *redisclient.RedisClient, cfg RateLimiterConfig) *RateLimiter {
	return &RateLimiter{
		redisClient:    *redisClient,
		redisLuaScript: redis.NewScript(utils.Constants.LuaScript),
		cfg:            cfg,
	}
}

func (rl *RateLimiter) Allow(ctx context.Context, key string) (bool, error) {
	redisKey := fmt.Sprintf("%s:%s", rl.cfg.KeyPrefix, key)

	now := time.Now().Unix()

	res, err := rl.redisLuaScript.Run(ctx, rl.redisClient.Rdb,
		[]string{redisKey},
		rl.cfg.Capacity,
		rl.cfg.RefillRate,
		now,
		1,
		int(rl.cfg.TTL.Seconds()),
	).Result()

	if err != nil {
		if rl.cfg.FailOpen {
			return true, nil
		}
		return false, err
	}

	allowed, ok := res.(int64)
	if !ok {
		return false, fmt.Errorf("unexpected result")
	}

	return allowed == 1, nil
}
