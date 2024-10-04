package middleware

import (
	"fmt"
	"simulation/internal/config"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(redis *config.RedisClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Cookies("session_id")
		if sessionID == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		userID, err := redis.Get(fmt.Sprintf("sess:%s", sessionID))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		c.Locals("userID", userID)
		return c.Next()
	}
}
