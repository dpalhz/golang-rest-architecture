package controller

import (
	"simulation/internal/controller/utils"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AnalyticController struct {
	Analyticervice *services.Analyticervice
}

func NewAnalyticController(Analyticervice *services.Analyticervice) *AnalyticController {
	return &AnalyticController{
		Analyticervice: Analyticervice,
	}
}

func (ac *AnalyticController) GetAllAnalyticHandler(c *fiber.Ctx) error {
	Analytic, err := ac.Analyticervice.GetAllAnalytic()
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, "Failed to retrieve Analytic", nil)
	}

	return utils.CreateResponse(c, fiber.StatusOK, true, "Analytic retrieved successfully", Analytic)
}