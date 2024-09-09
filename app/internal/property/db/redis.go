package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

type redisCache struct {
	cache  *redis.Client
	logger logging.Logger
}

type RedisCache interface {
}

func NewRedisCache(client *redis.Client, logger logging.Logger) RedisCache {
	return &redisCache{
		cache:  client,
		logger: logger,
	}
}
