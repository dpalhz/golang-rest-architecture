package routes

import (
	"simulation/internal/controller"
	"simulation/internal/repository"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AnalyticRoutes(api fiber.Router, db *gorm.DB) {
	analyticRepo := repository.NewAnalyticRepository(db)
	Analyticervice := services.NewAnalyticervice(analyticRepo)
	analyticController := controller.NewAnalyticController(Analyticervice)

	analyticRoutes := api.Group("/analytic")
	analyticRoutes.Get("/", analyticController.GetAllAnalyticHandler)

	
}
