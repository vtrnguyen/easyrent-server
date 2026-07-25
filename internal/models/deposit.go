package models

type Deposit struct {
	BaseModel
	RentalRequestID string        `gorm:"type:char(36);not null"`
	Amount          float64       `gorm:"type:decimal(15,2);not null"`
	TransactionCode string        `gorm:"type:varchar(255)"`
	Status          string        `gorm:"type:enum('pending','paid','refunded','failed','canceled');default:'pending'"`
	RentalRequest   RentalRequest `gorm:"foreignKey:RentalRequestID"`
}
