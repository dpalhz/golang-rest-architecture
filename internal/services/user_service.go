package services

import (
	"errors"
	"simulation/internal/entity"
	"simulation/internal/model/converter"
	"simulation/internal/model/request"
	"simulation/internal/model/response"
	"simulation/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepo  *repository.UserRepository
	Validator *validator.Validate
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo:  userRepo,
		Validator: validator.New(),
	}
}

func (s *UserService) RegisterUser(dto *request.UserRegister) (*response.UserRegister, error) {
	if err := s.Validator.Struct(dto); err != nil {
		return nil, err 
	}

	hashedPassword, err := HashPassword(dto.Password)
	if err != nil {
		return nil, err 
	}

	user := converter.ConvertRegisterRequestToUser(dto)
	user.Password = hashedPassword

	if err := s.UserRepo.Create(user); err != nil {
		return nil, err
	}
	return converter.ConvertUserToRegisterResponse(user), nil
}

func (s *UserService) LoginUser(dto *request.UserLogin) (*response.UserLogin, error) {
	if err := s.Validator.Struct(dto); err != nil {
		return nil, err 
	}

	var user entity.User
	err := s.UserRepo.FindByUsername(&user, dto.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if err := ComparePasswords(user.Password, dto.Password); err != nil {
		return nil, errors.New("invalid username or password") 
	}
	return converter.ConvertUserToLoginResponse(&user), nil
}

func (s *UserService) GetUserByID(id int) (*entity.User, error) {
    var user entity.User
    err := s.UserRepo.FindById(&user, id)
    if err != nil {
        return nil, errors.New("user not found")
    }
    return &user, nil
}

func (s *UserService) UpdateUser(dto *request.UpdateUser, userID int) (*response.UserUpdate, error) {
	if err := s.Validator.Struct(dto); err != nil {
		return nil, err 
	}

    user, err := s.GetUserByID(userID)
    if err != nil {
        return nil, err
    }

	if dto.Name != nil  {
		user.Name = *dto.Name 
	}
	if dto.Email != nil {
		user.Email = *dto.Email 
	}
	if dto.Photo != nil {
		user.Photo = *dto.Photo 
	}

	if err := s.UserRepo.Update(user); err != nil {
		return nil, err
	}
	return converter.ConvertUserToUpdateResponse(user), nil
}

func (s *UserService) DeleteUser(userID int) error {
    user, err := s.GetUserByID(userID)
    if err != nil {
        return err
    }
    return s.UserRepo.SoftDelete(user)
}

func (s *UserService) GetUserProfile(userID int) (*response.UserDetails, error) {
	// Fetch user from repository based on ID
	user, err := s.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	profileResponse :=  converter.ConvertUserToDetailsResponse(user)
	return profileResponse, nil
}

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

