package response

import (
	"net/http"

	"cinema-booking-api/internal/domain/dto"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, dto.SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func BadRequest(c *gin.Context, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}

	c.JSON(http.StatusBadRequest, dto.ErrorResponse{
		Success: false,
		Message: message,
		Error:   errorMsg,
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
		Success: false,
		Message: message,
	})
}

func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, dto.ErrorResponse{
		Success: false,
		Message: message,
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, dto.ErrorResponse{
		Success: false,
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}

	c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
		Success: false,
		Message: message,
		Error:   errorMsg,
	})
}
