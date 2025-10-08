package constant

type UserRole string

const (
	UserRoleCustomer UserRole = "customer"
	UserRoleAdmin    UserRole = "admin"
	UserRoleStaff    UserRole = "staff"
)

type SeatType string

const (
	SeatTypeRegular SeatType = "regular"
	SeatTypeVIP     SeatType = "vip"
	SeatTypePremium SeatType = "premium"
)

type ScheduleStatus string

const (
	ScheduleStatusActive    ScheduleStatus = "active"
	ScheduleStatusCancelled ScheduleStatus = "cancelled"
	ScheduleStatusCompleted ScheduleStatus = "completed"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusPaid      BookingStatus = "paid"
	BookingStatusCancelled BookingStatus = "cancelled"
	BookingStatusRefunded  BookingStatus = "refunded"
)

type BookingSeatStatus string

const (
	BookingSeatStatusLocked    BookingSeatStatus = "locked"
	BookingSeatStatusSold      BookingSeatStatus = "sold"
	BookingSeatStatusAvailable BookingSeatStatus = "available"
)

type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusApproved  RefundStatus = "approved"
	RefundStatusRejected  RefundStatus = "rejected"
	RefundStatusCompleted RefundStatus = "completed"
)
