package utils

import (
	"simulation/internal/model"

	"github.com/gofiber/fiber/v2"
)

func CreateResponse(c *fiber.Ctx, errorCode int, success bool, message string, data interface{}) error {
	response := model.APIResponse{
		ErrorCode: errorCode,
		Success:   success,
		Message:   message,
		Data:      data,
	}
	return c.Status(errorCode).JSON(response)
}
