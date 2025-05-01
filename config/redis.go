package common

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

var rdb *RedisClient

// InitializeRedis initializes and returns a Redis client
func InitializeRedis(addr, password string, db int) *RedisClient {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	rdb = &RedisClient{
		Client: client,
		Ctx:    ctx,
	}
	return rdb
}

func GetRedisClient() *RedisClient {
	if rdb == nil {
		panic("Redis client not initialized")
	}
	return rdb
}

// Common operations
func (r *RedisClient) Set(key string, value interface{}, ttl time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, ttl).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *RedisClient) Del(key string) error {
	return r.Client.Del(r.Ctx, key).Err()
}

func (r *RedisClient) Exists(key string) (bool, error) {
	n, err := r.Client.Exists(r.Ctx, key).Result()
	return n > 0, err
}
