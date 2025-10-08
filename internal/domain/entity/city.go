package entity

import (
	"time"

	"github.com/google/uuid"
)

type City struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Code      string    `gorm:"type:varchar(10);uniqueIndex;not null" json:"code"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (City) TableName() string {
	return "cities"
}
