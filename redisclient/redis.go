package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/truongle2004/service-context/logger"
)

var (
	RedisClient *redis.Client
	redisOnce   sync.Once
)

// InitRedis initializes the Redis client using singleton pattern
func InitRedis(addr, password string, db int) error {
	var initErr error

	redisOnce.Do(func() {
		ctx := context.Background()
		RedisClient = redis.NewClient(&redis.Options{
			Password: password,
			Addr:     addr,
			DB:       db,
		})

		if err := RedisClient.Ping(ctx).Err(); err != nil {
			initErr = err
		} else {
			logger.Info("Redis client initialized")
		}
	})

	return initErr
}

func GetClient() *redis.Client {
	if RedisClient == nil {
		panic("Redis not initialized")
	}
	return RedisClient
}

func Set(c *gin.Context, key string, value any, expiration time.Duration) error {
	return RedisClient.Set(c.Request.Context(), key, value, expiration).Err()
}

func Get(c *gin.Context, key string) (string, error) {
	return RedisClient.Get(c.Request.Context(), key).Result()
}

func Delete(c *gin.Context, key string) error {
	return RedisClient.Del(c.Request.Context(), key).Err()
}

func Exists(c *gin.Context, key string) (bool, error) {
	n, err := RedisClient.Exists(c.Request.Context(), key).Result()
	return n > 0, err
}

func SetJSON(c *gin.Context, key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return RedisClient.Set(c.Request.Context(), key, data, expiration).Err()
}

func GetJSON(c *gin.Context, key string, dest any) error {
	data, err := RedisClient.Get(c.Request.Context(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

func MatchingKey(c *gin.Context, key string) (bool, error) {
	pattern := fmt.Sprintf("*%s*", key)

	keys, err := RedisClient.Keys(c.Request.Context(), pattern).Result()
	if err != nil {
		return false, err
	}

	return len(keys) > 0, nil
}
