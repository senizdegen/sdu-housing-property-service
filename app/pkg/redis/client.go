package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/senizdegen/sdu-housing/property-service/pkg/logging"
)

func NewClient(addr, password string, logger logging.Logger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
