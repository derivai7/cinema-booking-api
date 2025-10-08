package entity

import (
	"cinema-booking-api/internal/constant"
	"time"

	"github.com/google/uuid"
)

type BookingSeat struct {
	ID         uuid.UUID                  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BookingID  uuid.UUID                  `gorm:"type:uuid;not null;index" json:"booking_id"`
	SeatID     uuid.UUID                  `gorm:"type:uuid;not null;index" json:"seat_id"`
	ScheduleID uuid.UUID                  `gorm:"type:uuid;not null;index" json:"schedule_id"`
	Status     constant.BookingSeatStatus `gorm:"type:booking_seat_status_enum;default:'locked';index" json:"status"`
	LockedAt   *time.Time                 `gorm:"index" json:"locked_at"`
	LockedBy   *uuid.UUID                 `gorm:"type:uuid" json:"locked_by"`
	CreatedAt  time.Time                  `gorm:"autoCreateTime" json:"created_at"`
}

func (BookingSeat) TableName() string {
	return "booking_seats"
}
