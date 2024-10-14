package services

import (
	"errors"
	"simulation/internal/entity"
	"simulation/internal/model/converter"
	"simulation/internal/model/request"
	"simulation/internal/model/response"
	"simulation/internal/repository"

	"gorm.io/gorm"
)

type Analyticervice struct {
	AnalyticRepo *repository.AnalyticRepository
}

func NewAnalyticervice(analyticRepo *repository.AnalyticRepository) *Analyticervice {
	return &Analyticervice{
		AnalyticRepo: analyticRepo,
	}
}

func (as *Analyticervice) GetAllAnalytic() ([]*entity.Analytic, error) {
	Analytic, err := as.AnalyticRepo.GetAllAnalytic()
	if err != nil {
		return nil, err
	}
	return Analytic, nil
}

func (s *Analyticervice) GetAnalyticByID(id uint) (*response.Analytic, error) {
	var analytic entity.Analytic
	err := s.AnalyticRepo.FindById(&analytic, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("analytic data not found")
		}
		return nil, err
	}
	return converter.ConvertAnalyticToAnalyticResponse(&analytic), nil
}

func (s *Analyticervice) CreateAnalytic(dto *request.CreateAnalytic) (*response.Analytic, error) {
	analytic := converter.ConvertCreateAnalyticRequestToAnalytic(dto)

	err := s.AnalyticRepo.Create(analytic)
	if err != nil {
		return nil, err
	}
	return converter.ConvertAnalyticToAnalyticResponse(analytic), nil
}

func (s *Analyticervice) UpdateAnalytic(id uint, dto *request.UpdateAnalytic) (*response.Analytic, error) {
	var analytic entity.Analytic
	err := s.AnalyticRepo.FindById(&analytic, id)
	if err != nil {
		return nil, err
	}
	converter.ConvertUpdateAnalyticRequestToAnalytic(dto, &analytic)

	err = s.AnalyticRepo.Update(&analytic)
	if err != nil {
		return nil, err
	}
	return converter.ConvertAnalyticToAnalyticResponse(&analytic), nil
}

func (s *Analyticervice) DeleteAnalytic(analyticId uint) error {
	var analytic entity.Analytic
	err := s.AnalyticRepo.FindById(&analytic, analyticId)
	if err != nil {
		return err
	}
	return s.AnalyticRepo.Delete(&analytic)
}
