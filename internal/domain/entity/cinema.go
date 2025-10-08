package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cinema struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CityID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"city_id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Address   string         `gorm:"type:text" json:"address"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	City City `gorm:"foreignKey:CityID;constraint:OnDelete:CASCADE" json:"city,omitempty"`
}

func (Cinema) TableName() string {
	return "cinemas"
}
