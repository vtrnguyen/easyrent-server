package models

import "time"

type PostComment struct {
	BaseModel
	UserID          string `gorm:"type:char(36);not null"`
	PostID          string `gorm:"type:char(36);not null"`
	ParentCommentID *string
	Content         string `gorm:"type:text;not null"`
	CreatedAt       time.Time
	User            User         `gorm:"foreignKey:UserID"`
	Post            Post         `gorm:"foreignKey:PostID"`
	ParentComment   *PostComment `gorm:"foreignKey:ParentCommentID"`
}
