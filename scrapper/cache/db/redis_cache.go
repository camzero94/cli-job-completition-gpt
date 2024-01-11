package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ttl   time.Duration
}

func NewRedisCache(c *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{
		client: c,
		ttl: ttl,
	}
}

func (c *RedisCache) Get(key string) (string, bool) {
	ctx := context.Background()
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (c *RedisCache) Set(key string, val string) error {
	ctx := context.Background()
	_, err := c.client.Set(ctx, key, val, c.ttl).Result()
	return err
}

func (c *RedisCache) Remove(key string) error {
	ctx := context.Background()
	_, err := c.client.Del(ctx, key).Result()
	return err
}
