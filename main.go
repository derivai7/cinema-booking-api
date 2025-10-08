package main

import (
	"log"

	"cinema-booking-api/internal/config"
	"cinema-booking-api/internal/http/handler"
	"cinema-booking-api/internal/http/route"
	"cinema-booking-api/internal/pkg/database"
	"cinema-booking-api/internal/pkg/jwt"
	"cinema-booking-api/internal/repository"
	"cinema-booking-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := database.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	jwtService := jwt.NewJWTService(cfg.JWT.Secret, cfg.JWT.ExpiredHours)

	userRepo := repository.NewUserRepository(db)
	scheduleRepo := repository.NewScheduleRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, jwtService)
	scheduleUsecase := usecase.NewScheduleUsecase(scheduleRepo)

	authHandler := handler.NewAuthHandler(authUsecase)
	scheduleHandler := handler.NewScheduleHandler(scheduleUsecase)

	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	route.SetupRoutes(router, &route.Config{
		AuthHandler:     authHandler,
		ScheduleHandler: scheduleHandler,
		JWTService:      jwtService,
	})

	log.Printf("Server starting on port %s", cfg.App.Port)
	if err = router.Run(":" + cfg.App.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
