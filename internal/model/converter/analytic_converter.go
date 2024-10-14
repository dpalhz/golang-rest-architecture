package converter

import (
	"simulation/internal/entity"
	"simulation/internal/model/request"
	"simulation/internal/model/response"
)

func ConvertCreateAnalyticRequestToAnalytic(dto *request.CreateAnalytic) *entity.Analytic {
	return &entity.Analytic{
		BlogID:    dto.BlogID,
		UserID:    dto.UserID,
		DeviceType: dto.DeviceType,
		Browser:    dto.Browser,
		IP:         dto.IP,
		Referrer:   dto.Referrer,
	}
}

func ConvertUpdateAnalyticRequestToAnalytic(dto *request.UpdateAnalytic, analytic *entity.Analytic) {
	if dto.BlogID != 0 {
		analytic.BlogID = dto.BlogID
	}
	if dto.UserID != 0 {
		analytic.UserID = dto.UserID
	}
	if dto.DeviceType != "" {
		analytic.DeviceType = dto.DeviceType
	}
	if dto.Browser != "" {
		analytic.Browser = dto.Browser
	}
	if dto.IP != "" {
		analytic.IP = dto.IP
	}
	if dto.Referrer != "" {
		analytic.Referrer = dto.Referrer
	}
}

func ConvertAnalyticToAnalyticResponse(analytic *entity.Analytic) *response.Analytic {
	return &response.Analytic{
		ID:        analytic.ID,
		BlogID:    analytic.BlogID,
		UserID:    analytic.UserID,
		DeviceType: analytic.DeviceType,
		Browser:    analytic.Browser,
		IP:         analytic.IP,
		Referrer:   analytic.Referrer,
		CreatedAt:  analytic.CreatedAt,
	}
}