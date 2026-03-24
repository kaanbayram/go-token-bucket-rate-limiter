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

Import `ratelimiter` (and `redisclient` if the bundled helper is used). Configuration and keys are supplied by the host application — see package docs and `server` for a runnable sketch.

**Packages:** `ratelimiter` — limiter API; `redisclient` — Redis helper; `server` — example entrypoint. `utils` is internal (Lua).
