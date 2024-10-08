package repository

import (
	"simulation/internal/entity"

	"gorm.io/gorm"
)

type BlogRepository struct {
	Repository[entity.Blog] 
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		Repository: Repository[entity.Blog]{DB: db},
	}
}

func (r *BlogRepository) GetBlogsPaginated(offset int, limit int) ([]entity.Blog, error) {
	var blogs []entity.Blog
	result := r.DB.Offset(offset).Limit(limit).
		Preload("Tags").    
		Preload("Categories"). 
		Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

func (r *BlogRepository) CountBlogs() (int64, error) {
	var count int64
	result := r.DB.Model(&entity.Blog{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}


func (r *BlogRepository) FindByIdAndAdminId(blogId uint, adminId uint) (*entity.Blog, error) {
	var blog entity.Blog
	result := r.DB.Where("id = ? AND admin_id = ?", blogId, adminId).Take(&blog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blog, nil
}

func (r *BlogRepository) FilterBlogsByAnyCategoriesAndTags(categories []uint, tags []uint) ([]entity.Blog, error) {
	var blogs []entity.Blog

	if len(categories) == 0 && len(tags) == 0 {
		result := r.DB.Preload("Categories").Preload("Tags").Find(&blogs)
		if result.Error != nil {
			return nil, result.Error
		}
		return blogs, nil
	}

	query := r.DB.Joins("JOIN blog_categories ON blogs.id = blog_categories.blog_id").
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id")


	if len(categories) > 0 {
		query = query.Where("blog_categories.category_id IN (?)", categories)
	}

	if len(tags) > 0 {
		query = query.Where("blog_tags.tag_id IN (?)", tags)
	}

	result := query.Preload("Categories").Preload("Tags").Find(&blogs)

	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
}


func (r *BlogRepository) FilterBlogsByAllCategoriesAndTags(categories []uint, tags []uint) ([]entity.Blog, error) {
	var blogs []entity.Blog

	// Start the query
	query := r.DB.Model(&entity.Blog{}).
		Joins("JOIN blog_categories ON blogs.id = blog_categories.blog_id").
		Joins("JOIN blog_tags ON blogs.id = blog_tags.blog_id").
		Group("blogs.id") 

	
	if len(categories) > 0 {
		query = query.Where("blog_categories.category_id IN (?)", categories)
	}

	if len(tags) > 0 {
		query = query.Where("blog_tags.tag_id IN (?)", tags)
	}

	if len(categories) > 0 {
		query = query.Having("COUNT(DISTINCT blog_categories.category_id) = ?", len(categories))
	}
	if len(tags) > 0 {
		query = query.Having("COUNT(DISTINCT blog_tags.tag_id) = ?", len(tags))
	}

	query = query.Preload("Categories").Preload("Tags")

	if err := query.Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}

