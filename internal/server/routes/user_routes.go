// server/routes/user_routes.go
package routes

import (
	"simulation/internal/controller"
	"simulation/internal/repository"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserRoutes defines all user-related routes

func UserRoutes(api fiber.Router, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userRoutes := api.Group("/user")
	userRoutes.Get("/:id", userController.GetUserProfileHandler)  
	userRoutes.Post("/register", userController.RegisterUserHandler)  
	userRoutes.Post("/login", userController.LoginUserHandler)        
	userRoutes.Put("/:id", userController.UpdateUserHandler)          
	userRoutes.Delete("/:id", userController.DeleteUserHandler)       
}
