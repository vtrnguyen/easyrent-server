package models

import "time"

type Appointment struct {
	BaseModel
	RentalRequestID string        `gorm:"type:char(36);not null"`
	AppointmentTime time.Time     `gorm:"not null"`
	Note            string        `gorm:"type:text"`
	Status          string        `gorm:"type:enum('pending','confirmed','completed','canceled','missed');default:'pending'"`
	RentalRequest   RentalRequest `gorm:"foreignKey:RentalRequestID"`
}
