package models

import (
	"time"
)

type User struct {
	BaseModel
	Status              string `gorm:"type:enum('active','inactive');default:'active'"`
	EmailVerified       bool   `gorm:"default:false"`
	LastLoginAt         *time.Time
	Account             Account              `gorm:"foreignKey:UserID"`
	UserInfo            UserInfo             `gorm:"foreignKey:UserID"`
	Properties          []Property           `gorm:"foreignKey:OwnerID"`
	Posts               []Post               `gorm:"foreignKey:AuthorID"`
	WantedPosts         []WantedPost         `gorm:"foreignKey:AuthorID"`
	RentalRequests      []RentalRequest      `gorm:"foreignKey:TenantID"`
	LandlordContracts   []Contract           `gorm:"foreignKey:LandlordID"`
	TenantContracts     []Contract           `gorm:"foreignKey:TenantID"`
	Notifications       []Notification       `gorm:"foreignKey:UserID"`
	AIConversations     []AIConversation     `gorm:"foreignKey:UserID"`
	SentMessages        []Message            `gorm:"foreignKey:SenderID"`
	ConversationMembers []ConversationMember `gorm:"foreignKey:UserID"`
	Reports             []Report             `gorm:"foreignKey:ReporterID"`
	PostLikes           []PostLike           `gorm:"foreignKey:UserID"`
	PostFavorites       []PostFavorite       `gorm:"foreignKey:UserID"`
	PostComments        []PostComment        `gorm:"foreignKey:UserID"`
}
