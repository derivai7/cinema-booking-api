package repository

import (
	"cinema-booking-api/internal/domain/entity"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	Create(schedule *entity.Schedule) error
	FindAll() ([]entity.Schedule, error)
	FindByID(id uuid.UUID) (*entity.Schedule, error)
	Update(schedule *entity.Schedule) error
	Delete(id uuid.UUID) error
	ExistsByMovieStudioDateTime(movieID, studioID uuid.UUID, showDate, showTime string) (bool, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(schedule *entity.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) FindAll() ([]entity.Schedule, error) {
	var schedules []entity.Schedule

	err := r.db.
		Preload("Movie").
		Preload("Studio.Cinema.City").
		Order("show_date DESC, show_time DESC").
		Find(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *scheduleRepository) FindByID(id uuid.UUID) (*entity.Schedule, error) {
	var schedule entity.Schedule

	err := r.db.
		Preload("Movie").
		Preload("Studio.Cinema.City").
		Where("id = ?", id).
		First(&schedule).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("schedule not found")
		}
		return nil, err
	}

	return &schedule, nil
}

func (r *scheduleRepository) Update(schedule *entity.Schedule) error {
	return r.db.Save(schedule).Error
}

func (r *scheduleRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entity.Schedule{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("schedule not found")
	}

	return nil
}

func (r *scheduleRepository) ExistsByMovieStudioDateTime(movieID, studioID uuid.UUID, showDate, showTime string) (bool, error) {
	var count int64

	err := r.db.Model(&entity.Schedule{}).
		Where("movie_id = ? AND studio_id = ? AND show_date = ? AND show_time = ?",
			movieID, studioID, showDate, showTime).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
