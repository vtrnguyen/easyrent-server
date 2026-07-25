package models

type Conversation struct {
	BaseModel
	PropertyID string               `gorm:"type:char(36);not null"`
	Property   Property             `gorm:"foreignKey:PropertyID"`
	Members    []ConversationMember `gorm:"foreignKey:ConversationID"`
	Messages   []Message            `gorm:"foreignKey:ConversationID"`
}
