package server

import (
	"simulation/internal/server/routes"

	"github.com/gofiber/fiber/v2"
)

// RegisterFiberRoutes defines all routes for the application
func (s *FiberServer) RegisterFiberRoutes() {

	api := s.App.Group("/api/v1")
	api.Get("/", s.HelloWorldHandler)
	api.Get("/db/health", s.healthHandler)

	routes.UserRoutes(api, s.db.GetDB())  
	routes.AuthRoutes(api, s.db.GetDB(), &s.redis)	
}


func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
