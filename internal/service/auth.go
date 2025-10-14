package service

import (
	"context"
	"digital-queue-system/internal/model"
	"digital-queue-system/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user with this username or email already exists")
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(ctx context.Context, userIdentity, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsernameOrEmail(ctx, userIdentity)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":      "auth-service",
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Minute * 10).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		log.Error().Err(err).Str("jwtsecret", s.jwtSecret).Msg("failed to signed token string")
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Register(ctx context.Context, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate password")
		return err
	}

	user := &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateUser) {
			return ErrUserAlreadyExists
		}
		return err
	}

	return nil
}
