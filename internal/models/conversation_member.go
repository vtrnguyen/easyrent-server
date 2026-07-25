package models

import "time"

type ConversationMember struct {
	ConversationID string `gorm:"primaryKey"`
	UserID         string `gorm:"primaryKey"`
	JoinedAt       time.Time
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	User           User         `gorm:"foreignKey:UserID"`
}
