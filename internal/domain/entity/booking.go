package entity

import (
	"time"

	"cinema-booking-api/internal/constant"

	"github.com/google/uuid"
)

type Booking struct {
	ID            uuid.UUID              `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID        uuid.UUID              `gorm:"type:uuid;not null;index" json:"user_id"`
	ScheduleID    uuid.UUID              `gorm:"type:uuid;not null;index" json:"schedule_id"`
	BookingCode   string                 `gorm:"type:varchar(20);uniqueIndex;not null" json:"booking_code"`
	TotalPrice    float64                `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status        constant.BookingStatus `gorm:"type:booking_status_enum;default:'pending';index" json:"status"`
	PaymentMethod string                 `gorm:"type:varchar(50)" json:"payment_method"`
	PaidAt        *time.Time             `json:"paid_at"`
	CreatedAt     time.Time              `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time              `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Booking) TableName() string {
	return "bookings"
}
