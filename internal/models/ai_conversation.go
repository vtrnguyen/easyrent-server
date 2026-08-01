package models

type AIConversation struct {
	BaseModel
	UserID string `gorm:"type:char(36);not null"`
	Title  string `gorm:"type:varchar(255)"`

	User     User        `gorm:"foreignKey:UserID"`
	Messages []AIMessage `gorm:"foreignKey:ConversationID"`
}