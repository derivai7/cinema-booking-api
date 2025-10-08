package route

import (
	"cinema-booking-api/internal/constant"
	"cinema-booking-api/internal/http/handler"
	"cinema-booking-api/internal/http/middleware"
	"cinema-booking-api/internal/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Config struct {
	AuthHandler     *handler.AuthHandler
	ScheduleHandler *handler.ScheduleHandler
	JWTService      *jwt.JWTService
}

func SetupRoutes(router *gin.Engine, cfg *Config) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", cfg.AuthHandler.Login)
		}

		schedules := api.Group("/schedules")
		schedules.Use(middleware.AuthMiddleware(cfg.JWTService))
		{
			schedules.GET("", cfg.ScheduleHandler.GetAllSchedules)
			schedules.GET("/:id", cfg.ScheduleHandler.GetScheduleByID)

			schedules.POST("",
				middleware.RequireRole(constant.UserRoleAdmin, constant.UserRoleStaff),
				cfg.ScheduleHandler.CreateSchedule)

			schedules.PUT("/:id",
				middleware.RequireRole(constant.UserRoleAdmin, constant.UserRoleStaff),
				cfg.ScheduleHandler.UpdateSchedule)

			schedules.DELETE("/:id",
				middleware.RequireRole(constant.UserRoleAdmin, constant.UserRoleStaff),
				cfg.ScheduleHandler.DeleteSchedule)
		}
	}
}
