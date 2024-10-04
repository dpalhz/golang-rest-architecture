package controller

import (
	"simulation/internal/controller/utils"
	"simulation/internal/model/request"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var dto request.UserLogin
	if err := c.BodyParser(&dto); err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid request payload", nil)
	}

	userLoginResp, err := ctrl.AuthService.UserAuthentication(&dto)
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusUnauthorized, false, err.Error(), nil)
	}

	ctrl.AuthService.SetSessionCookie(c, userLoginResp.SessionID)

	return utils.CreateResponse(c, fiber.StatusOK, true, "Login successful", userLoginResp)
}


func (ctrl *AuthController) Logout(c *fiber.Ctx) error {
	sessionID := c.Cookies("session_id")

	if err := ctrl.AuthService.LogoutUser(sessionID); err != nil {
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, "Failed to logout", nil)
	}

	ctrl.AuthService.DeleteSessionCookie(c, sessionID)

	return utils.CreateResponse(c, fiber.StatusOK, true, "Logout successful", nil)
}
