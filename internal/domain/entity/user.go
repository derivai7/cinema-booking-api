package entity

import (
	"time"

	"cinema-booking-api/internal/constant"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email     string            `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string            `gorm:"type:varchar(255);not null" json:"-"`
	FullName  string            `gorm:"type:varchar(255);not null" json:"full_name"`
	Phone     string            `gorm:"type:varchar(20)" json:"phone"`
	Role      constant.UserRole `gorm:"type:user_role_enum;default:'customer'" json:"role"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
