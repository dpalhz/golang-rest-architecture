package repository

import (
	"simulation/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User] 
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: Repository[entity.User]{DB: db},
	}
}

func (r *UserRepository) FindByUsername(user *entity.User, username string) error {
	return r.DB.Where("username = ?", username).Take(user).Error
}

func (r *UserRepository) FindByEmail(user *entity.User, email string) error {
	return r.DB.Where("email = ?", email).Take(user).Error
}
