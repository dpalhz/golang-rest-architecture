package controller

import (
	"simulation/internal/controller/utils"
	"simulation/internal/model/request"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) UserRegisterHandler(c *fiber.Ctx) error {
	dto := new(request.UserRegister)

	if err := c.BodyParser(dto); err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid input", nil)
	}

	response, err := uc.UserService.RegisterUser(dto)
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return utils.CreateResponse(c, fiber.StatusCreated, true, "User registered successfully", response)
}

func (uc *UserController) UserLoginHandler(c *fiber.Ctx) error {
	dto := new(request.UserLogin)

	// Parse body
	if err := c.BodyParser(dto); err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid input", nil)
	}

	response, err := uc.UserService.LoginUser(dto)
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusUnauthorized, false, err.Error(), nil)
	}

	return utils.CreateResponse(c, fiber.StatusOK, true, "Login successful", response)
}

func (uc *UserController) UserUpdateHandler(c *fiber.Ctx) error {
	dto := new(request.UpdateUser)

	userID, err := c.ParamsInt("id")
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid user ID", nil)
	}

	if err := c.BodyParser(dto); err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid input", nil)
	}

	response, err := uc.UserService.UpdateUser(dto, userID)
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return utils.CreateResponse(c, fiber.StatusOK, true, "User updated successfully", response)
}

func (uc *UserController) UserDeleteHandler(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid user ID", nil)
	}

	if err := uc.UserService.DeleteUser(userID); err != nil {
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return utils.CreateResponse(c, fiber.StatusOK, true, "User deleted successfully", nil)
}


func (uc *UserController) ProfileUserHandler(c *fiber.Ctx) error {
	// Dapatkan user ID dari parameter URL
	userID, err := c.ParamsInt("id")
	if err != nil || userID <= 0 {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid user ID", nil)
	}

	profile, err := uc.UserService.GetUserProfile(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return utils.CreateResponse(c, fiber.StatusNotFound, false, err.Error(), nil)
		}
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, err.Error(), nil)
	}

	return utils.CreateResponse(c, fiber.StatusOK, true, "User profile retrieved successfully", profile)
}