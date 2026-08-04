package models

type Account struct {
	BaseModel
	UserID      string `gorm:"type:char(36);not null"`
	Role        string `gorm:"type:enum('admin', 'tenant', 'landlord');default:'tenant'"`
	Email       string
	PhoneNumber string
	Password    string
}
