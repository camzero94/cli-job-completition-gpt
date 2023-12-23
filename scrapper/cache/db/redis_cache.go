
package db

import (
	"github.com/go-redis/redis/v8"
	"context"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(c *redis.Client) *RedisCache {
	return &RedisCache{
		client: c,
	}	
}

func (c *RedisCache) Get(key string) (string, bool){
	ctx := context.Background()
	val, err := c.client.Get(ctx,key).Result()
	if err != nil{
		return "",false
	}
	return val, true
}

func (c *RedisCache) Set(key string, val string) error{
	ctx := context.Background()
	_, err := c.client.Set(ctx, key,val,0).Result()
	return err
}

func (c *RedisCache) Remove(key string )error{
	ctx := context.Background()
	_, err := c.client.Del(ctx, key).Result()
	return err
}




