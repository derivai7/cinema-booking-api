package usecase

import (
	"errors"
	"fmt"
	"time"

	"cinema-booking-api/internal/domain/dto"
	"cinema-booking-api/internal/domain/entity"
	"cinema-booking-api/internal/repository"

	"github.com/google/uuid"
)

type ScheduleUsecase interface {
	CreateSchedule(req dto.CreateScheduleRequest) (*dto.ScheduleResponse, error)
	GetAllSchedules() ([]dto.ScheduleResponse, error)
	GetScheduleByID(id string) (*dto.ScheduleResponse, error)
	UpdateSchedule(id string, req dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error)
	DeleteSchedule(id string) error
}

type scheduleUsecase struct {
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleUsecase(scheduleRepo repository.ScheduleRepository) ScheduleUsecase {
	return &scheduleUsecase{
		scheduleRepo: scheduleRepo,
	}
}

func (u *scheduleUsecase) CreateSchedule(req dto.CreateScheduleRequest) (*dto.ScheduleResponse, error) {
	movieID, err := uuid.Parse(req.MovieID)
	if err != nil {
		return nil, errors.New("invalid movie ID format")
	}

	studioID, err := uuid.Parse(req.StudioID)
	if err != nil {
		return nil, errors.New("invalid studio ID format")
	}

	showDate, err := time.Parse("2006-01-02", req.ShowDate)
	if err != nil {
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	if _, err := time.Parse("15:04", req.ShowTime); err != nil {
		return nil, errors.New("invalid time format, expected HH:MM (24-hour)")
	}

	if req.Price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	exists, err := u.scheduleRepo.ExistsByMovieStudioDateTime(movieID, studioID, req.ShowDate, req.ShowTime)
	if err != nil {
		return nil, fmt.Errorf("failed to check schedule existence: %w", err)
	}
	if exists {
		return nil, errors.New("schedule already exists for this movie, studio, date, and time")
	}

	schedule := &entity.Schedule{
		MovieID:  movieID,
		StudioID: studioID,
		ShowDate: showDate,
		ShowTime: req.ShowTime,
		Price:    req.Price,
		Status:   "active",
	}

	if err := u.scheduleRepo.Create(schedule); err != nil {
		return nil, fmt.Errorf("failed to create schedule: %w", err)
	}

	createdSchedule, err := u.scheduleRepo.FindByID(schedule.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve created schedule: %w", err)
	}

	return u.toScheduleResponse(createdSchedule), nil
}

func (u *scheduleUsecase) GetAllSchedules() ([]dto.ScheduleResponse, error) {
	schedules, err := u.scheduleRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve schedules: %w", err)
	}

	responses := make([]dto.ScheduleResponse, 0, len(schedules))
	for i := range schedules {
		responses = append(responses, *u.toScheduleResponse(&schedules[i]))
	}

	return responses, nil
}

func (u *scheduleUsecase) GetScheduleByID(id string) (*dto.ScheduleResponse, error) {
	scheduleID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid schedule ID format")
	}

	schedule, err := u.scheduleRepo.FindByID(scheduleID)
	if err != nil {
		return nil, err
	}

	return u.toScheduleResponse(schedule), nil
}

func (u *scheduleUsecase) UpdateSchedule(id string, req dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error) {
	scheduleID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid schedule ID format")
	}

	schedule, err := u.scheduleRepo.FindByID(scheduleID)
	if err != nil {
		return nil, err
	}

	if req.MovieID != "" {
		movieID, err := uuid.Parse(req.MovieID)
		if err != nil {
			return nil, errors.New("invalid movie ID format")
		}
		schedule.MovieID = movieID
	}

	if req.StudioID != "" {
		studioID, err := uuid.Parse(req.StudioID)
		if err != nil {
			return nil, errors.New("invalid studio ID format")
		}
		schedule.StudioID = studioID
	}

	if req.ShowDate != "" {
		showDate, err := time.Parse("2006-01-02", req.ShowDate)
		if err != nil {
			return nil, errors.New("invalid date format, expected YYYY-MM-DD")
		}
		schedule.ShowDate = showDate
	}

	if req.ShowTime != "" {
		if _, err := time.Parse("15:04", req.ShowTime); err != nil {
			return nil, errors.New("invalid time format, expected HH:MM (24-hour)")
		}
		schedule.ShowTime = req.ShowTime
	}

	if req.Price > 0 {
		schedule.Price = req.Price
	}

	if req.Status != "" {
		if req.Status != "active" && req.Status != "cancelled" && req.Status != "completed" {
			return nil, errors.New("invalid status, must be: active, cancelled, or completed")
		}
		schedule.Status = req.Status
	}

	if err := u.scheduleRepo.Update(schedule); err != nil {
		return nil, fmt.Errorf("failed to update schedule: %w", err)
	}

	updatedSchedule, err := u.scheduleRepo.FindByID(scheduleID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated schedule: %w", err)
	}

	return u.toScheduleResponse(updatedSchedule), nil
}

func (u *scheduleUsecase) DeleteSchedule(id string) error {
	scheduleID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid schedule ID format")
	}

	_, err = u.scheduleRepo.FindByID(scheduleID)
	if err != nil {
		return err
	}

	if err := u.scheduleRepo.Delete(scheduleID); err != nil {
		return fmt.Errorf("failed to delete schedule: %w", err)
	}

	return nil
}

func (u *scheduleUsecase) toScheduleResponse(schedule *entity.Schedule) *dto.ScheduleResponse {
	response := &dto.ScheduleResponse{
		ID:        schedule.ID,
		MovieID:   schedule.MovieID,
		StudioID:  schedule.StudioID,
		ShowDate:  schedule.ShowDate,
		ShowTime:  schedule.ShowTime,
		Price:     schedule.Price,
		Status:    schedule.Status,
		CreatedAt: schedule.CreatedAt,
		UpdatedAt: schedule.UpdatedAt,
	}

	if schedule.Movie.ID != uuid.Nil {
		response.Movie = &dto.MovieInfo{
			ID:          schedule.Movie.ID,
			Title:       schedule.Movie.Title,
			Duration:    schedule.Movie.Duration,
			Genre:       schedule.Movie.Genre,
			Rating:      schedule.Movie.Rating,
			ReleaseDate: schedule.Movie.ReleaseDate,
		}
	}

	if schedule.Studio.ID != uuid.Nil {
		studioInfo := &dto.StudioInfo{
			ID:         schedule.Studio.ID,
			Name:       schedule.Studio.Name,
			TotalSeats: schedule.Studio.TotalSeats,
		}

		if schedule.Studio.Cinema.ID != uuid.Nil {
			cinemaInfo := dto.CinemaInfo{
				ID:      schedule.Studio.Cinema.ID,
				Name:    schedule.Studio.Cinema.Name,
				Address: schedule.Studio.Cinema.Address,
			}

			if schedule.Studio.Cinema.City.ID != uuid.Nil {
				cinemaInfo.City = dto.CityInfo{
					ID:   schedule.Studio.Cinema.City.ID,
					Name: schedule.Studio.Cinema.City.Name,
					Code: schedule.Studio.Cinema.City.Code,
				}
			}

			studioInfo.Cinema = cinemaInfo
		}

		response.Studio = studioInfo
	}

	return response
}
