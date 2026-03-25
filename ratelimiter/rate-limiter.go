package ratelimiter

import (
	"context"
	"fmt"
	"go-token-bucket-rate-limiter/utils"
	"strings"
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
	redisClient    redis.Cmdable
	redisLuaScript *redis.Script
	cfg            RateLimiterConfig
}

func New(redisClient redis.Cmdable, cfg RateLimiterConfig) (*RateLimiter, error) {
	if redisClient == nil {
		return nil, fmt.Errorf("redis client is required")
	}
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return &RateLimiter{
		redisClient:    redisClient,
		redisLuaScript: redis.NewScript(utils.Constants.LuaScript),
		cfg:            cfg,
	}, nil
}

func (rl *RateLimiter) AllowContext(ctx context.Context, key string) (bool, error) {
	redisKey := fmt.Sprintf("%s:%s", rl.cfg.KeyPrefix, key)

	now := time.Now().Unix()

	res, err := rl.redisLuaScript.Run(ctx, rl.redisClient,
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

func (rl *RateLimiter) Allow(key string) (bool, error) {
	return rl.AllowContext(context.Background(), key)
}

func validateConfig(cfg RateLimiterConfig) error {
	if cfg.Capacity <= 0 {
		return fmt.Errorf("capacity must be greater than 0")
	}
	if cfg.RefillRate <= 0 {
		return fmt.Errorf("refill rate must be greater than 0")
	}
	if cfg.TTL <= 0 {
		return fmt.Errorf("ttl must be greater than 0")
	}
	if strings.TrimSpace(cfg.KeyPrefix) == "" {
		return fmt.Errorf("key prefix is required")
	}

	return nil
}
