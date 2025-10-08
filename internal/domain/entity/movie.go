package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Duration    int            `gorm:"not null" json:"duration"`
	Genre       string         `gorm:"type:varchar(100)" json:"genre"`
	Rating      string         `gorm:"type:varchar(10)" json:"rating"`
	PosterURL   string         `gorm:"type:varchar(500)" json:"poster_url"`
	ReleaseDate time.Time      `gorm:"type:date" json:"release_date"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Movie) TableName() string {
	return "movies"
}
