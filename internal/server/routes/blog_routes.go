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

func BlogRoutes(api fiber.Router, db *gorm.DB, rd *config.RedisClient) {
	blogRepo := repository.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)
	cacheMiddleware := middleware.NewCacheMiddleware(rd.Client)

	blogRoutes := api.Group("/blogs")
	blogRoutes.Get("/", cacheMiddleware.Cache, blogController.GetBlogsPaginatedHandler)
	blogRoutes.Get("/filter", blogController.FilterBlogsHandler)

}



