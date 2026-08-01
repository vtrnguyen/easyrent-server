package models

type AIMessage struct {
	BaseModel

	ConversationID string `gorm:"type:char(36);not null"`
	Role           string `gorm:"type:enum('system','user','assistant');not null"`
	Content        string `gorm:"type:longtext;not null"`

	Conversation AIConversation `gorm:"foreignKey:ConversationID"`
}