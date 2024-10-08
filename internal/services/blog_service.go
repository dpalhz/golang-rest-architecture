package services

import (
	"fmt"
	"simulation/internal/entity"
	"simulation/internal/model/converter"
	"simulation/internal/model/request"
	"simulation/internal/model/response"
	"simulation/internal/repository"

	"gopkg.in/go-playground/validator.v9"
)

// BlogService handles the business logic for the Blog entity
type BlogService struct {
	BlogRepo *repository.BlogRepository
	Validator *validator.Validate
}

func NewBlogService(repo *repository.BlogRepository) *BlogService {
	return &BlogService{
		BlogRepo: repo,
		Validator: validator.New()}
}

func (s *BlogService) PaginateBlogs(dto *request.Blogs) (*response.Blogs, error) {

	if err := s.Validator.Struct(dto); err != nil {
		return nil, err 
	}

	page := dto.Page
	limit := dto.Limit
	
	offset := (page - 1) * limit

	fmt.Println(offset, limit)

	blogs, err := s.BlogRepo.GetBlogsPaginated(offset, limit)
	if err != nil {
		return nil, err
	}

	totalCount, err := s.BlogRepo.CountBlogs()
	if err != nil {
		return nil, err
	}
	return converter.ConvertBlogsToBlogsResponse(blogs, totalCount, page, limit), nil
}



func (s *BlogService) CreateBlog(dto *request.CreateBlog) (*response.Blog, error) {
	// Validate request
	err := s.Validator.Struct(dto)
	if err != nil {
		return nil, err
	}
	blog := converter.ConvertCreateBlogRequestToBlog(dto)

	err = s.BlogRepo.Create(blog)
	if err != nil {
		return nil, err
	}
	return converter.ConvertBlogToBlogResponse(blog), nil
}


func (s *BlogService) UpdateBlog(id uint, dto *request.UpdateBlog) (*response.Blog, error) {
	// Find existing blog
	var blog entity.Blog
	err := s.BlogRepo.FindById(&blog, id)
	if err != nil {
		return nil, err
	}
	converter.ConvertUpdateBlogRequestToBlog(dto, &blog)

	err = s.BlogRepo.Update(&blog)
	if err != nil {
		return nil, err
	}
	return converter.ConvertBlogToBlogResponse(&blog), nil
}

func (s *BlogService) DeleteBlog(blogId uint, adminId uint) error {
	blog, err := s.BlogRepo.FindByIdAndAdminId(blogId, adminId)
	if err != nil {
		return err 
	}
	return s.BlogRepo.Delete(blog)
}

func (s *BlogService) GetBlogByID(id uint) (*response.Blog, error) {
	var blog entity.Blog
	err := s.BlogRepo.FindById(&blog, id)
	if err != nil {
		return nil, err
	}
	return converter.ConvertBlogToBlogResponse(&blog), nil
}

func (s *BlogService) FilterBlogs(dto *request.BlogFilter) ([]*response.Blog, error) {
    var blogs []entity.Blog
    var err error

    if dto.MatchAll {
        blogs, err = s.BlogRepo.FilterBlogsByAllCategoriesAndTags(dto.Categories, dto.Tags)
    } else {
        blogs, err = s.BlogRepo.FilterBlogsByAnyCategoriesAndTags(dto.Categories, dto.Tags)
    }

    if err != nil {
        return nil, err
    }

    var blogResponses []*response.Blog
    for _, blog := range blogs {
        blogResponses = append(blogResponses, converter.ConvertBlogToBlogResponse(&blog))
    }

    return blogResponses, nil
}
