package server

import (
	"github.com/gofiber/fiber/v2"

	"simulation/internal/config"
)

type FiberServer struct {
	*fiber.App

	db config.Gorm
	redis config.RedisClient
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "simulation",
			AppName:      "simulation",
		}),

		db: config.NewGorm(),
		redis: *config.NewRedisClient(),
	}

	return server
}
