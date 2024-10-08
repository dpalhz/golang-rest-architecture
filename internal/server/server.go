package server

import (
	"simulation/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	*fiber.App

	DB config.Gorm
	Redis config.RedisClient
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			// Prefork: true,
			ServerHeader: "simulation",
			AppName:      "simulation",
		}),

		DB: config.NewGorm(),
		Redis: *config.NewRedisClient(),
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins: "https://yourdomain.com, https://anotherdomain.com",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	

	return server
}
