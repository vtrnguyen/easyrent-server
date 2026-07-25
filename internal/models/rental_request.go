package models

type RentalRequest struct {
	BaseModel
	TenantID     string        `gorm:"type:char(36);not null"`
	PropertyID   string        `gorm:"type:char(36);not null"`
	Message      string        `gorm:"type:text"`
	Status       string        `gorm:"type:enum('pending','accepted','rejected','canceled');default:'pending'"`
	Tenant       User          `gorm:"foreignKey:TenantID"`
	Property     Property      `gorm:"foreignKey:PropertyID"`
	Appointments []Appointment `gorm:"foreignKey:RentalRequestID"`
	Deposits     []Deposit     `gorm:"foreignKey:RentalRequestID"`
	Payments     []Payment     `gorm:"foreignKey:RentalRequestID"`
}
