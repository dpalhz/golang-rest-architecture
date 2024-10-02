package server

import (
	"github.com/gofiber/fiber/v2"

	"simulation/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "simulation",
			AppName:      "simulation",
		}),

		db: database.New(),
	}

	return server
}
