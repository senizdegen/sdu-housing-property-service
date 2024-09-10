package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

type redisCache struct {
	client *redis.Client
	logger logging.Logger
}

type RedisCache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}

func NewRedisCache(client *redis.Client, logger logging.Logger) RedisCache {
	return &redisCache{
		client: client,
		logger: logger,
	}
}

func (r *redisCache) Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *redisCache) Get(key string) (string, error) {
	ctx := context.Background()
	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		//Key does not exist
		return "", nil
	} else if err != nil {
		return "", err
	}
	return result, nil
}
