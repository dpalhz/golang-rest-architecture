package routes

import (
	"simulation/internal/controller"
	"simulation/internal/repository"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(api fiber.Router, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userRoutes := api.Group("/user")
	userRoutes.Get("/:id", userController.ProfileUserHandler)  
	userRoutes.Post("/register", userController.UserRegisterHandler)  
	userRoutes.Post("/login", userController.UserLoginHandler)        
	userRoutes.Put("/:id", userController.UserUpdateHandler)          
	userRoutes.Delete("/:id", userController.UserDeleteHandler)       
}
