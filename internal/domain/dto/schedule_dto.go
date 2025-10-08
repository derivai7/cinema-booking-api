package dto

import (
	"cinema-booking-api/internal/constant"
	"time"

	"github.com/google/uuid"
)

type CreateScheduleRequest struct {
	MovieID  string  `json:"movie_id" binding:"required,uuid4"`
	StudioID string  `json:"studio_id" binding:"required,uuid4"`
	ShowDate string  `json:"show_date" binding:"required"`
	ShowTime string  `json:"show_time" binding:"required"`
	Price    float64 `json:"price" binding:"required,min=0"`
}

type UpdateScheduleRequest struct {
	MovieID  string                  `json:"movie_id" binding:"omitempty,uuid4"`
	StudioID string                  `json:"studio_id" binding:"omitempty,uuid4"`
	ShowDate string                  `json:"show_date" binding:"omitempty"`
	ShowTime string                  `json:"show_time" binding:"omitempty"`
	Price    float64                 `json:"price" binding:"omitempty,min=0"`
	Status   constant.ScheduleStatus `json:"status" binding:"omitempty,oneof=active cancelled completed"`
}

type ScheduleResponse struct {
	ID        uuid.UUID               `json:"id"`
	MovieID   uuid.UUID               `json:"movie_id"`
	StudioID  uuid.UUID               `json:"studio_id"`
	ShowDate  time.Time               `json:"show_date"`
	ShowTime  string                  `json:"show_time"`
	Price     float64                 `json:"price"`
	Status    constant.ScheduleStatus `json:"status"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	Movie     *MovieInfo              `json:"movie,omitempty"`
	Studio    *StudioInfo             `json:"studio,omitempty"`
}

type MovieInfo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Duration    int       `json:"duration"`
	Genre       string    `json:"genre"`
	Rating      string    `json:"rating"`
	ReleaseDate time.Time `json:"release_date"`
}

type StudioInfo struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	TotalSeats int        `json:"total_seats"`
	Cinema     CinemaInfo `json:"cinema"`
}

type CinemaInfo struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	City    CityInfo  `json:"city"`
}

type CityInfo struct {
	ID   uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}
