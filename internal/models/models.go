package models

import (
	"time"
)

// User represents the user table in the database
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FirstName    string    `gorm:"size:50;not null" json:"first_name"`
	LastName     string    `gorm:"size:50;not null" json:"last_name"`
	Email        string    `gorm:"size:100;unique;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	PhoneNumber  string    `gorm:"size:15" json:"phone_number"`
	Gender       string    `gorm:"size:10" json:"gender"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Barber represents the barber table in the database
type Barber struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	FirstName       string    `gorm:"size:50;not null" json:"first_name"`
	LastName        string    `gorm:"size:50;not null" json:"last_name"`
	ProfilePicture  string    `gorm:"size:255" json:"profile_picture"`
	Bio             string    `gorm:"type:text" json:"bio"`
	SocialMediaLink string    `gorm:"size:255" json:"social_media_link"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Shop represents the shop table in the database
type Shop struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Address   string    `gorm:"size:255" json:"address"`
	State     string    `gorm:"size:50" json:"state"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Service represents the service table in the database
type Service struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BarberID    uint      `gorm:"not null" json:"barber_id"`
	ShopID      uint      `gorm:"not null" json:"shop_id"`
	ServiceName string    `gorm:"size:100;not null" json:"service_name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Duration    int       `gorm:"not null" json:"duration"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Booking represents the booking table in the database
type Booking struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	ServiceID       uint      `gorm:"not null" json:"service_id"`
	BarberID        uint      `gorm:"not null" json:"barber_id"`
	ShopID          uint      `gorm:"not null" json:"shop_id"`
	AppointmentTime time.Time `gorm:"not null" json:"appointment_time"`
	Status          string    `gorm:"size:20;default:'Pending'" json:"status"`
	PaymentStatus   string    `gorm:"size:20;default:'Pending'" json:"payment_status"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Payment represents the payment table in the database
type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	BookingID     uint      `gorm:"not null" json:"booking_id"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod string    `gorm:"size:50;not null" json:"payment_method"`
	PaymentStatus string    `gorm:"size:20;default:'Pending'" json:"payment_status"`
	PaymentDate   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"payment_date"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// PromoCode represents the promocode table in the database
type PromoCode struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Code               string    `gorm:"size:50;unique;not null" json:"code"`
	DiscountPercentage float64   `gorm:"type:decimal(5,2);not null" json:"discount_percentage"`
	ExpirationDate     time.Time `gorm:"not null" json:"expiration_date"`
	UsageLimit         int       `gorm:"not null" json:"usage_limit"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Admin represents the admin table in the database
type Admin struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"size:50;unique;not null" json:"username"`
	Email        string    `gorm:"size:100;unique;not null" json:"email"`
	Phone        string    `gorm:"size:15" json:"phone"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Role         string    `gorm:"size:20;default:'Admin'" json:"role"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
