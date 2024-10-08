package controller

import (
	"fmt"
	"simulation/internal/controller/utils"
	"simulation/internal/model/request"
	"simulation/internal/services"

	"github.com/gofiber/fiber/v2"
)


type BlogController struct {
	blogService *services.BlogService
}


func NewBlogController(blogService *services.BlogService) *BlogController {
	return &BlogController{
		blogService: blogService}
}


func (bc *BlogController) GetBlogsPaginatedHandler(c *fiber.Ctx) error {
	fmt.Println("Entering GetBlogsPaginatedHandler")
	dto := new(request.Blogs)

	if err := c.BodyParser(dto); err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid input", nil)
	}

	if dto.Page < 1 {
		dto.Page = 1
	}

	if dto.Limit <= 0 {
		dto.Limit = 8
	}

	fmt.Printf("Fetching blogs for page %d with limit %d\n", dto.Page, dto.Limit)

	response, err := bc.blogService.PaginateBlogs(dto)
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, "Failed to retrieve blogs", nil)
	}

	return utils.CreateResponse(c, fiber.StatusOK, true, "Blogs retrieved successfully", response)
}


func (bc *BlogController) FilterBlogsHandler(c *fiber.Ctx) error {
	dto := new(request.BlogFilter)

	if err := c.BodyParser(dto); err != nil {
		return utils.CreateResponse(c, fiber.StatusBadRequest, false, "Invalid input", nil)
	}

	blogs, err := bc.blogService.FilterBlogs(dto)
	if err != nil {
		return utils.CreateResponse(c, fiber.StatusInternalServerError, false, "Failed to filter blogs", nil)
	}

	return utils.CreateResponse(c, fiber.StatusOK, true, "Blogs filtered successfully", blogs)
}