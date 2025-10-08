package entity

import (
	"time"

	"cinema-booking-api/internal/constant"

	"github.com/google/uuid"
)

type Refund struct {
	ID           uuid.UUID             `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BookingID    uuid.UUID             `gorm:"type:uuid;not null;index" json:"booking_id"`
	Amount       float64               `gorm:"type:decimal(10,2);not null" json:"amount"`
	Reason       string                `gorm:"type:text" json:"reason"`
	RefundMethod string                `gorm:"type:varchar(50)" json:"refund_method"`
	Status       constant.RefundStatus `gorm:"type:refund_status_enum;default:'pending';index" json:"status"`
	ProcessedAt  *time.Time            `json:"processed_at"`
	CreatedAt    time.Time             `gorm:"autoCreateTime" json:"created_at"`
}

func (Refund) TableName() string {
	return "refunds"
}
