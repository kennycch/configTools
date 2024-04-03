package redis

import "github.com/go-redis/redis"

const (
	LogoutTokens = "logout_tokens"
)

var (
	RD *redis.Client
)
