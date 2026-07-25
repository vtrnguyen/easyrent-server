package models

type Message struct {
	BaseModel
	ConversationID string       `gorm:"type:char(36);not null"`
	SenderID       string       `gorm:"type:char(36);not null"`
	MessageType    string       `gorm:"type:enum('text','image','video','file','system');not null"`
	MediaURL       string       `gorm:"type:text"`
	Content        string       `gorm:"type:longtext"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	Sender         User         `gorm:"foreignKey:SenderID"`
}
