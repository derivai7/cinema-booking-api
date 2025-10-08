package handler

import (
	"cinema-booking-api/internal/domain/dto"
	"cinema-booking-api/internal/pkg/response"
	"cinema-booking-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err)
		return
	}

	result, err := h.authUsecase.Login(req)
	if err != nil {
		response.BadRequest(c, "Login failed", err)
		return
	}

	response.Success(c, "Login successful", result)
}
