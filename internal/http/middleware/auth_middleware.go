package middleware

import (
	"strings"
	
	"cinema-booking-api/internal/constant"
	"cinema-booking-api/internal/pkg/jwt"
	"cinema-booking-api/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *jwt.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "Authorization header is required")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "Invalid authorization format")
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			response.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func RequireRole(allowedRoles ...constant.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			response.Forbidden(c, "Access denied: role not found")
			c.Abort()
			return
		}

		roleAllowed := false
		for _, allowedRole := range allowedRoles {
			allowedRoleStr := string(allowedRole)
			if role == allowedRoleStr {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			response.Forbidden(c, "Access denied: you don't have permission to access this resource")
			c.Abort()
			return
		}

		c.Next()
	}
}
