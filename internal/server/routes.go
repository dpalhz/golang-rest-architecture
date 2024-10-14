package server

import (
	"simulation/internal/server/routes"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {

	api := s.App.Group("/api/v1")
	api.Get("/", s.HelloWorldHandler)
	api.Get("/db/health", s.healthHandler)

	routes.UserRoutes(api, s.DB.GetDB())  
	routes.AnalyticRoutes(api, s.DB2.GetDB())
	routes.BlogRoutes(api, s.DB.GetDB(), &s.Redis)  
	routes.AuthRoutes(api, s.DB.GetDB(), &s.Redis)	
}


func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.DB.Health())
}
