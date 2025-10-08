package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Studio struct {
	ID         uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CinemaID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"cinema_id"`
	Name       string         `gorm:"type:varchar(50);not null" json:"name"`
	TotalSeats int            `gorm:"not null" json:"total_seats"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	Cinema Cinema `gorm:"foreignKey:CinemaID;constraint:OnDelete:CASCADE" json:"cinema,omitempty"`
}

func (Studio) TableName() string {
	return "studios"
}
