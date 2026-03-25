# go-token-bucket-rate-limiter

Distributed token-bucket rate limiting for Go using Redis and a Lua script for atomic updates.

**Status:** Work in progress — API may change before a stable release.

## Requirements

- Go 1.25+
- Redis

## Install

```bash
go get go-token-bucket-rate-limiter@latest
```

Replace the module path with the published path (e.g. `github.com/<user>/go-token-bucket-rate-limiter`) once `go.mod` matches the repo.

Depends on [go-redis v9](https://github.com/redis/go-redis).

## Quick usage

```go
ctx := context.Background()
rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

if err := rdb.Ping(ctx).Err(); err != nil {
	log.Fatal(err)
}

limiter, err := ratelimiter.New(rdb, ratelimiter.RateLimiterConfig{
	Capacity:   10,
	RefillRate: 2,
	TTL:        60 * time.Second,
	KeyPrefix:  "rate_limit",
	FailOpen:   true,
})
if err != nil {
	log.Fatal(err)
}

allowed, err := limiter.Allow("127.0.0.1")
if err != nil {
	log.Fatal(err)
}
fmt.Println("allowed:", allowed)
```

Import `ratelimiter` directly; `redisclient` is optional helper code. Configuration and keys are supplied by the host application.

Run the example app:

```bash
go run ./cmd/example
```

**Packages:** `ratelimiter` — limiter API; `redisclient` — Redis helper; `cmd/example` — runnable example app. `utils` is internal (Lua).
