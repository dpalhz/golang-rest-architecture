package repository

import (
	"simulation/internal/entity"

	"gorm.io/gorm"
)

type AnalyticRepository struct {
	Repository[entity.Analytic]
}

func NewAnalyticRepository(db *gorm.DB) *AnalyticRepository {
	return &AnalyticRepository{
		Repository: Repository[entity.Analytic]{DB: db},
	}
}

func (r *AnalyticRepository) GetAllAnalytic() ([]*entity.Analytic, error) {
	var Analytic []*entity.Analytic
	err := r.DB.Find(&Analytic).Error
	if err != nil {
		return nil, err
	}
	return Analytic, nil
}