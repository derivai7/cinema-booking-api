package usecase

import (
	"errors"

	"cinema-booking-api/internal/domain/dto"
	"cinema-booking-api/internal/pkg/jwt"
	"cinema-booking-api/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authUsecase struct {
	userRepo   repository.UserRepository
	jwtService *jwt.JWTService
}

func NewAuthUsecase(userRepo repository.UserRepository, jwtService *jwt.JWTService) AuthUsecase {
	return &authUsecase{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (u *authUsecase) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := u.jwtService.GenerateToken(user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	response := &dto.LoginResponse{
		Token: token,
		User: dto.UserData{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
			Phone:    user.Phone,
			Role:     user.Role,
		},
	}

	return response, nil
}
