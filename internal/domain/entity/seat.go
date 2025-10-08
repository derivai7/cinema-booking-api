package entity

import (
	"time"

	"cinema-booking-api/internal/constant"

	"github.com/google/uuid"
)

type Seat struct {
	ID         uuid.UUID         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	StudioID   uuid.UUID         `gorm:"type:uuid;not null;index" json:"studio_id"`
	SeatNumber string            `gorm:"type:varchar(10);not null" json:"seat_number"`
	Row        string            `gorm:"type:varchar(5);not null" json:"row"`
	SeatType   constant.SeatType `gorm:"type:seat_type_enum;default:'regular'" json:"seat_type"`
	CreatedAt  time.Time         `gorm:"autoCreateTime" json:"created_at"`
}

func (Seat) TableName() string {
	return "seats"
}
