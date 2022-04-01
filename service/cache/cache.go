package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/readreceipt/api/config"
)

var r *redis.Client

func Init() error {
	r = redis.NewClient(&redis.Options{
		Addr:     config.RedisURL(),
		Password: "",
		DB:       config.RedisDatabase(),
	})

	return r.Ping(context.TODO()).Err()
}

func Store(key string, value any, ttl time.Duration) error {
	return r.Set(context.TODO(), key, value, ttl).Err()
}

func Exists(key string) (bool, error) {
	exists, err := r.Exists(context.TODO(), key).Result()

	return exists == 1, err
}
