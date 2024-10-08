package converter

import (
	"simulation/internal/entity"
	"simulation/internal/model/request"
	"simulation/internal/model/response"
)
func ConvertUserToRegisterResponse(user *entity.User) *response.UserRegister {
	return &response.UserRegister{ // Kembalikan pointer ke UserRegister
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ConvertRegisterRequestToUser(dto *request.UserRegister) *entity.User {
	return &entity.User{
		Name:     dto.Name,
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func ConvertUserToLoginResponse(user *entity.User) *response.UserLogin {
	return &response.UserLogin{
		ID:   user.ID,
		Name: user.Name,
	}
}

func ConvertUserToDetailsResponse(user *entity.User) *response.UserDetails {
	return &response.UserDetails{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Photo:    user.Photo,
	}
}

func ConvertUserToUpdateResponse(user *entity.User) *response.UserUpdate {
	return &response.UserUpdate{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Photo: user.Photo,
	}
}
