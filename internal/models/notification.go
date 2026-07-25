package models

type Notification struct {
	BaseModel
	UserID  string `gorm:"type:char(36);not null"`
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:text;not null"`
	Type    string `gorm:"type:enum('post_liked', 'new_comment', 'new_message', 'appointment_created', 'appointment_confirmed', 'deposit_success', 'payment_success', 'rental_request_created', 'post_approved', 'system');not null"`
	IsRead  bool   `gorm:"default:false"`
	User    User   `gorm:"foreignKey:UserID"`
}
