package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
    Client *redis.Client
}

var redisInstance *RedisClient

func NewRedisClient() *RedisClient {
    if redisInstance != nil {
        return redisInstance
    }

    redisAddr := os.Getenv("REDIS_ADDR")
    redisPassword := os.Getenv("REDIS_PASSWORD")
    redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
    if err != nil {
        redisDB = 0
    }

    client := redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: redisPassword,
        DB:       redisDB,       
    })

    pong, err := client.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }
    fmt.Println("Connected to Redis:", pong)

    redisInstance = &RedisClient{
        Client: client,
    }
    return redisInstance
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
    err := r.Client.Set(ctx, key, value, expiration).Err()
    if err != nil {
        return fmt.Errorf("failed to set key %s: %v", key, err)
    }
    return nil
}

func (r *RedisClient) Get(key string) (string, error) {
    value, err := r.Client.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("key %s does not exist", key)
    } else if err != nil {
        return "", fmt.Errorf("could not get key %s: %v", key, err)
    }
    return value, nil
}

func (r *RedisClient) Delete(key string) error {
    err := r.Client.Del(ctx, key).Err()
    if err != nil {
        return fmt.Errorf("failed to delete key %s: %v", key, err)
    }
    return nil
}

func (r *RedisClient) Close() error {
    return r.Client.Close()
}
