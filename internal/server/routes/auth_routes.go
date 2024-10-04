// server/routes/user_routes.go
package routes

import (
	"simulation/internal/config"
	"simulation/internal/controller"
	"simulation/internal/middleware"
	"simulation/internal/repository"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(api fiber.Router, db *gorm.DB, rd *config.RedisClient) {
	authMiddleware := middleware.SessionMiddleware(rd)
	authRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(authRepo, rd)
	authController := controller.NewAuthController(authService)

	authRoutes := api.Group("/auth")
	authRoutes.Post("/login", authController.Login)
	authRoutes.Post("/logout", authMiddleware, authController.Logout)
}
