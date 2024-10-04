package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"simulation/internal/config"
	"simulation/internal/entity"
	"simulation/internal/model/converter"
	"simulation/internal/model/request"
	"simulation/internal/model/response"
	"simulation/internal/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

type AuthService struct {
	UserRepo  *repository.UserRepository
	Validator *validator.Validate
	redisClient *config.RedisClient
}

func NewAuthService(userRepo *repository.UserRepository, redisClient *config.RedisClient) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		Validator: validator.New(),
		redisClient: redisClient,
	}
}

func (s *AuthService) UserAuthentication(dto *request.UserLogin) (*response.UserLogin, error) {
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
	sessionID := generateSessionID() 
	sessionKey := "sess:" + sessionID
	err = s.redisClient.Set(sessionKey, user.ID, 10*time.Minute) 
	if err != nil {
		return nil, errors.New("failed to create session")
	}

	loginResponse := converter.ConvertUserToLoginResponse(&user)
	loginResponse.SessionID = sessionID 
	return loginResponse, nil
}


func (s *AuthService) LogoutUser(sessionID string) error {
	sessionKey := "sess:" + sessionID
	
	if err := s.redisClient.Delete(sessionKey); err != nil {
		return errors.New("failed to delete session")
	}
	return nil
}


func (s *AuthService) SetSessionCookie(c *fiber.Ctx, sessionID string) {
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    sessionID, 
		Expires:  time.Now().Add(10 * time.Minute), 
	
	})
}

func (s *AuthService) DeleteSessionCookie(c *fiber.Ctx, sessionID string) {
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
	
	})
}

func generateSessionID() string {
	randomBytes := make([]byte, 8) 
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Handle error
		return "error_generating_session_id"
	}
	sessionID := fmt.Sprintf("%x-%d", randomBytes, time.Now().UnixNano())
	return sessionID
}

func ComparePasswords(hashedPassword, plainPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
