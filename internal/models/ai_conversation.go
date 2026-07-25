package models

type AIConversation struct {
	BaseModel
	UserID   string `gorm:"type:char(36);not null"`
	Question string `gorm:"type:longtext;not null"`
	Answer   string `gorm:"type:longtext;not null"`
	User     User   `gorm:"foreignKey:UserID"`
}
