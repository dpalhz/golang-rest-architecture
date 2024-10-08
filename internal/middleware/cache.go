// server/middleware/cache.go
package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type CacheMiddleware struct {
	RedisClient *redis.Client
}

func NewCacheMiddleware(redisClient *redis.Client) *CacheMiddleware {
	return &CacheMiddleware{RedisClient: redisClient}
}
func (m *CacheMiddleware) Cache(c *fiber.Ctx) error {
    page := c.Query("page", "1")
    limit := c.Query("limit", "8")

    cacheKey := fmt.Sprintf("blogs:page:%s:limit:%s", page, limit)

    // Check if the response is in the cache
    cachedResponse, err := m.RedisClient.Get(c.Context(), cacheKey).Result()
    if err == nil {
        fmt.Println("Cache hit: serving response from Redis")
        c.Set("Content-Type", "application/json")
        return c.Send([]byte(cachedResponse))
    }

    fmt.Println("Cache miss: proceeding to next handler")
    
    err = c.Next()
    if err != nil {
        return err
    }

    responseBody := c.Response().Body()
    if err := m.RedisClient.Set(c.Context(), cacheKey, responseBody, time.Minute*10).Err(); err != nil {
        fmt.Println("Error setting response in cache:", err)
    }

    fmt.Println("Response stored in cache")
    c.Set("Content-Type", "application/json")
    return c.Send(responseBody)
}
