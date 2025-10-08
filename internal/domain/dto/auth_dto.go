package dto

import (
	"cinema-booking-api/internal/constant"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}

type UserData struct {
	ID       uuid.UUID         `json:"id"`
	Email    string            `json:"email"`
	FullName string            `json:"full_name"`
	Phone    string            `json:"phone"`
	Role     constant.UserRole `json:"role"`
}
