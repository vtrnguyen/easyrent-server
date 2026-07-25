package models

import "time"

type Bill struct {
	BaseModel
	ContractID          string    `gorm:"type:char(36);not null"`
	BillingPeriod       time.Time `gorm:"type:date;not null"`
	RoomFee             float64   `gorm:"type:decimal(15,2);not null"`
	ElectricityOldIndex int
	ElectricityNewIndex int
	ElectricityUsage    int
	ElectricityFee      float64 `gorm:"type:decimal(15,2)"`
	WaterOldIndex       int
	WaterNewIndex       int
	WaterUsage          int
	WaterFee            float64   `gorm:"type:decimal(15,2)"`
	ServiceFee          float64   `gorm:"type:decimal(15,2)"`
	DiscountAmount      float64   `gorm:"type:decimal(15,2)"`
	TotalAmount         float64   `gorm:"type:decimal(15,2);not null"`
	DueDate             time.Time `gorm:"type:date;not null"`
	PaidAt              *time.Time
	Note                string   `gorm:"type:text"`
	Status              string   `gorm:"type:enum('draft','pending','paid','overdue','canceled');default:'draft'"`
	Contract            Contract `gorm:"foreignKey:ContractID"`
}
