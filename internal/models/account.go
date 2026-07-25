package models

type Account struct {
	BaseModel
	UserID      string
	Role        string `gorm:"type:enum('admin', 'tenant', 'landlord');default:'tenant'"`
	Email       string
	PhoneNumber string
	Password    string
}
