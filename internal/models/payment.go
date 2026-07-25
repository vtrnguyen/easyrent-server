package models

import "time"

type Payment struct {
	BaseModel
	RentalRequestID string  `gorm:"type:char(36);not null"`
	Amount          float64 `gorm:"type:decimal(15,2);not null"`
	PaymentMethod   string  `gorm:"type:enum('cash','bank_transfer','momo');not null"`
	PaymentStatus   string  `gorm:"type:enum('pending','paid','failed','refunded','canceled');default:'pending'"`
	PaidAt          *time.Time
	RentalRequest   RentalRequest `gorm:"foreignKey:RentalRequestID"`
}
