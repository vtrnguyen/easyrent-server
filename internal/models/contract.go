package models

import "time"

type Contract struct {
	BaseModel
	LandlordID      string    `gorm:"type:char(36);not null"`
	TenantID        string    `gorm:"type:char(36);not null"`
	PropertyID      string    `gorm:"type:char(36);not null"`
	StartDate       time.Time `gorm:"type:date;not null"`
	EndDate         time.Time `gorm:"type:date;not null"`
	RentPrice       float64   `gorm:"type:decimal(15,2);not null"`
	DepositAmount   float64   `gorm:"type:decimal(15,2);not null"`
	ContractFileURL string    `gorm:"type:text"`
	Status          string    `gorm:"type:enum('draft','active','expired','terminated');default:'draft'"`
	Landlord        User      `gorm:"foreignKey:LandlordID"`
	Tenant          User      `gorm:"foreignKey:TenantID"`
	Property        Property  `gorm:"foreignKey:PropertyID"`
	Bills           []Bill    `gorm:"foreignKey:ContractID"`
}
