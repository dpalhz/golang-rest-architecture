package converter

import (
	"simulation/internal/entity"
	"simulation/internal/model/request"
	"simulation/internal/model/response"
)

func ConvertBlogToBlogDetailResponse(blog *entity.Blog) *response.BlogDetail {
	categories := make([]string, len(blog.Categories))
	for i, category := range blog.Categories {
		categories[i] = category.Name
	}

	tags := make([]string, len(blog.Tags))
	for i, tag := range blog.Tags {
		tags[i] = tag.Name
	}
	return &response.BlogDetail{
		ID:         blog.ID,
		Title:      blog.Title,
		Content:    blog.Content,
		ReadTime:   blog.ReadTime,
		IsBlocked:  blog.IsBlocked,
		Categories: categories,
	
	}
}


func ConvertBlogsToBlogsResponse(blogs []entity.Blog, totalCount int64, page, limit int) *response.Blogs {
	var blogDetails []response.BlogDetail

	for _, blog := range blogs {
		blogDetails = append(blogDetails, *ConvertBlogToBlogDetailResponse(&blog))
	}
	return &response.Blogs{
		Blogs:      blogDetails,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   limit,
	}
}

func ConvertCreateBlogRequestToBlog(dto *request.CreateBlog) *entity.Blog {
	return &entity.Blog{
		Title:     dto.Title,
		Content:   dto.Content,
		ReadTime:  dto.ReadTime,
		IsBlocked: dto.IsBlocked,
		AdminID:   dto.AdminID,

	}
}

func ConvertUpdateBlogRequestToBlog(dto *request.UpdateBlog, blog *entity.Blog) {
	blog.Title = dto.Title
	blog.Content = dto.Content
	blog.ReadTime = dto.ReadTime
	blog.IsBlocked = dto.IsBlocked
}

func ConvertBlogToBlogResponse(blog *entity.Blog) *response.Blog {
	var categories []string
	for _, category := range blog.Categories {
		categories = append(categories, category.Name)
	}

	var tags []string
	for _, tag := range blog.Tags {
		tags = append(tags, tag.Name)
	}

	return &response.Blog{
		ID:         blog.ID,
		Title:      blog.Title,
		Content:    blog.Content,
		ReadTime:   blog.ReadTime,
		IsBlocked:  blog.IsBlocked,
		AdminID:    blog.AdminID,
		Categories: categories,
		Tags:       tags,
		CreatedAt:  blog.CreatedAt,
		UpdatedAt:  blog.UpdatedAt,
	}
}

