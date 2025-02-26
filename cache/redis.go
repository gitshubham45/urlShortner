package cache

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func RedisInit() *redis.Client {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	return client
}

func CacheURL(redisClient *redis.Client, shortUrl string, longUrl string) error {
	err := redisClient.Set(ctx, shortUrl, longUrl, 0).Err() // No expiration time
	if err != nil {
		return err
	}
	return nil
}

func GetCachedURL(redisClient *redis.Client, shortUrl string) (string, error) {
	longUrl, err := redisClient.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		return "", nil // URL not found in cache
	} else if err != nil {
		return "", err
	}
	return longUrl, nil
}
