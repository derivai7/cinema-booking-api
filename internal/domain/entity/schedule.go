package entity

import (
	"time"
	
	"cinema-booking-api/internal/constant"

	"github.com/google/uuid"
)

type Schedule struct {
	ID        uuid.UUID               `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	MovieID   uuid.UUID               `gorm:"type:uuid;not null;index" json:"movie_id"`
	StudioID  uuid.UUID               `gorm:"type:uuid;not null;index" json:"studio_id"`
	ShowDate  time.Time               `gorm:"type:date;not null;index" json:"show_date"`
	ShowTime  string                  `gorm:"type:time;not null" json:"show_time"`
	Price     float64                 `gorm:"type:decimal(10,2);not null" json:"price"`
	Status    constant.ScheduleStatus `gorm:"type:schedule_status_enum;default:'active';index" json:"status"`
	CreatedAt time.Time               `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time               `gorm:"autoUpdateTime" json:"updated_at"`

	Movie  Movie  `gorm:"foreignKey:MovieID;constraint:OnDelete:CASCADE" json:"movie,omitempty"`
	Studio Studio `gorm:"foreignKey:StudioID;constraint:OnDelete:CASCADE" json:"studio,omitempty"`
}

func (Schedule) TableName() string {
	return "schedules"
}
